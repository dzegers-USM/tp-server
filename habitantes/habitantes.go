package main

import (
	// "context"

	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"

	// "io"
	"log"
	"math/rand"
	"net"

	// "os"
	"sync"
	"time"

	pb "github.com/dzegers-USM/tp-server/habitantes/misc"

	"google.golang.org/grpc"
)

type coordenadaAgua struct {
	X int32
	Y int32
}

type habitante struct {
	posX   int32
	posY   int32
	estado int32
	id     int32
}

type server struct {
	pb.UnimplementedServicioHabitantesServer
	habitantes []habitante
	mu         sync.Mutex // Protege el acceso concurrente a las variables
}

type autoInc struct {
	sync.Mutex
	id int32
}

func (a *autoInc) ID() (id int32) {
	a.Lock()
	defer a.Unlock()
	a.id++
	id = a.id

	return
}

var autoIncrement autoInc // global instance

/*
***	Set parametros iniciales para cada habitante y comunicarlos a cliente.
***	Cliente debe comunicar el numero de habitantes a setear en su request.
 */
func (s *server) InicializadorHabitantes(in *pb.InicializadorRequest, stream pb.ServicioHabitantes_InicializadorHabitantesServer) error {

	//Parametros
	const min_posX int32 = 0
	const max_posX int32 = 600
	const min_posY int32 = 0
	const max_posY int32 = 600
	const min_sed int32 = 35
	const max_sed int32 = 100

	//Protege el acceso concurrente a las variables mientras se inicializan.
	s.mu.Lock()
	defer s.mu.Unlock()

	//Set parametros iniciales para cada habitante y comunicarlos a cliente.
	s.habitantes = make([]habitante, in.GetNumHabitantes())
	for i := range s.habitantes {
		s.habitantes[i] = habitante{
			id:     autoIncrement.ID(),
			posX:   rand.Int31n(max_posX-min_posX+1) + min_posX,
			posY:   rand.Int31n(max_posY-min_posY+1) + min_posY,
			estado: rand.Int31n(max_sed-min_sed+1) + min_sed,
		}
	}
	fmt.Printf("Iniciando con n=%d habitantes", in.GetNumHabitantes())

	//Enviar stream de estado inicial a cliente
	response := &pb.InicializadorResponse{HabitantesInicial: formatHabitantesResponse(s.habitantes)}
	if err := stream.Send(response); err != nil {
		return err
	}

	return nil
}

// Convertir habitantes a []*pb.Habitante para poder enviarlo como respuesta.
func formatHabitantesResponse(habitantes []habitante) []*pb.Habitante {
	responseHabitantes := make([]*pb.Habitante, len(habitantes))
	for i, h := range habitantes {
		responseHabitantes[i] = &pb.Habitante{
			PosX:   h.posX,
			PosY:   h.posY,
			Estado: h.estado,
			Id:     h.id,
		}
	}
	return responseHabitantes
}

/*
***	Mueve a los habitantes de manera aleatorea cada 1 seg.
***	Cada 5 segundos actualiza el nivel de sed de los habitantes (resta 1 a cada uno) y se lo informa al cliente.
 */
func (s *server) ActualizarEstado(in *pb.EstadoRequest, stream pb.ServicioHabitantes_ActualizarEstadoServer) error {

	//Coordenadas awa
	coordenadas, err := readCoordenates("../../Civitas/coordenadas.csv")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	if len(coordenadas) == 0 {
		log.Fatalf("No se encontraron coordenadas en el archivo")
	}

	//Timers
	ticker_sed := time.NewTicker(1000 * time.Millisecond) //n(000) Segundos
	defer ticker_sed.Stop()

	ticker_movimiento := time.NewTicker(500 * time.Millisecond) //n(000) Segundos
	defer ticker_movimiento.Stop()

	ticker_reproduccion := time.NewTicker(10000 * time.Millisecond)
	defer ticker_reproduccion.Stop()

	for {
		select {
		case <-ticker_sed.C:
			s.mu.Lock()
			s.updateEstados()
			response := &pb.EstadoResponse{EstadoHabitante: formatHabitantesResponse(s.habitantes)}
			s.mu.Unlock()

			if err := stream.Send(response); err != nil {
				return err
			}
		case <-ticker_movimiento.C:
			s.mu.Lock()
			s.randomWalk()
			s.walkToWater(coordenadas)
			response := &pb.EstadoResponse{EstadoHabitante: formatHabitantesResponse(s.habitantes)}
			s.mu.Unlock()

			if err := stream.Send(response); err != nil {
				return err
			}
		case <-ticker_reproduccion.C:
			s.mu.Lock()
			s.reproduccion()
			response := &pb.EstadoResponse{EstadoHabitante: formatHabitantesResponse(s.habitantes)}
			s.mu.Unlock()

			if err := stream.Send(response); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}

/*
***	Disminuye estado de habitantes en -1 simulando sed.
 */
func (s *server) updateEstados() {
	for i := range s.habitantes {
		if s.habitantes[i].estado > 0 {
			s.habitantes[i].estado -= 1
		}
	}
}

func (s *server) reproduccion() {
	for i := range s.habitantes {
		if s.habitantes[i].estado > 110 {
			s.habitantes[i].estado -= 10
			s.habitantes = append(s.habitantes, habitante{
				id:     autoIncrement.ID(),
				posX:   s.habitantes[i].posX,
				posY:   s.habitantes[i].posY,
				estado: 100,
			})
		}
	}
}

/*
***	Setea estado de habitante en 100.
 */
// func (s *server) ConsumirRecurso(in *pb.EstadoRequest, stream pb.ServicioHabitantes_ConsumirRecursoServer) error {
// 	//Protege el acceso concurrente a las variables mientras se ejecuta la funcion.
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	var id_habitante int32 = in.GetRequest()

// 	if s.habitantes[id_habitante].estado > 0 {
// 		s.habitantes[id_habitante].estado += 100
// 	}

// 	response := &pb.EstadoResponse{Respuesta: s.habitantes}
// 	if err := stream.Send(response); err != nil {
// 		return err
// 	}
// 	return stream.Context().Err()
// }

func (s *server) randomWalk() {

	var umbralSed int32 = 30

	for i := range s.habitantes {

		if s.habitantes[i].estado >= umbralSed {
			//Coordenadas X
			if s.habitantes[i].posX <= 0 {
				s.habitantes[i].posX = int32(s.habitantes[i].posX) + rand.Int31n(2) //Suma 0 o 1
			} else if s.habitantes[i].posX >= 600 {
				s.habitantes[i].posX = int32(s.habitantes[i].posX) + rand.Int31n(2) - 1 //Suma 0 o -1
			} else {
				s.habitantes[i].posX = int32(s.habitantes[i].posX) + rand.Int31n(3) - 1 //Suma -1, 0 o -1
			}

			//Coordenadas Y
			if s.habitantes[i].posY <= 0 {
				s.habitantes[i].posY = int32(s.habitantes[i].posY) + rand.Int31n(2) //Suma 0 o 1
			} else if s.habitantes[i].posX >= 600 {
				s.habitantes[i].posY = int32(s.habitantes[i].posY) + rand.Int31n(2) - 1 //Suma 0 o -1
			} else {
				s.habitantes[i].posY = int32(s.habitantes[i].posY) + rand.Int31n(3) - 1 //Suma -1, 0 o -1
			}
		}
	}
}

/*
***	Lee coordenadas desde el archivo fileName y las retorna.
 */
func readCoordenates(fileName string) ([]coordenadaAgua, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	lineas, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer el archivo csv: %v", err)
	}

	var coordenadas []coordenadaAgua

	for _, line := range lineas {

		x, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("no se pudo convertir X a entero: %v", err)
		}

		y, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, fmt.Errorf("no se pudo convertir Y a entero: %v", err)
		}

		coordenadas = append(coordenadas, coordenadaAgua{
			X: 100*int32(x) + 50,
			Y: 100*int32(y) + 50,
		})
	}

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	var x, y, e int32
	// 	_, err := fmt.Sscanf(scanner.Text(), "%d;%d;%d", &x, &y, &e)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	coordenadas = append(coordenadas, coordenadaAgua{X: x*100 + 50, Y: y*100 + 50})
	// }
	// if err := scanner.Err(); err != nil {
	// 	return nil, err
	// }
	return coordenadas, nil
}

/*
***	Calcula la distancia entre dos coordenadas c1 y c2.
 */
func calcularDistancia(c1, c2 coordenadaAgua) float64 {
	dx := float64(c1.X - c2.X)
	dy := float64(c1.Y - c2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

/*
Calcula y retorna la coordenada de la casilla de agua mas cercana al habitante.
*/
func closestWater(origen habitante, coordenadas []coordenadaAgua) (coordenadaAgua, error) {
	var newOrigen coordenadaAgua
	newOrigen.X = origen.posX
	newOrigen.Y = origen.posY

	coordenadaMasCercana := coordenadas[0]
	distanciaMinima := calcularDistancia(newOrigen, coordenadaMasCercana)

	for _, coordenada := range coordenadas[1:] {
		distancia := calcularDistancia(newOrigen, coordenada)
		if distancia < distanciaMinima {
			coordenadaMasCercana = coordenada
			distanciaMinima = distancia
		}
	}

	return coordenadaMasCercana, nil
}

/*
***	Hace que habitante correspondiente a idHabitante de un paso en direccion a destino.
 */
// func (s *server) walkToWater(idHabitante int32, destino coordenadaAgua) {

// 	var paso int32 = 1 //Si se quiere poder cambiar la velocidad hay que modificar la funcion

// 	//Movimiento X
// 	if s.habitantes[idHabitante].posX < destino.X {
// 		s.habitantes[idHabitante].posX = int32(s.habitantes[idHabitante].posX) + paso
// 	} else if s.habitantes[idHabitante].posX > destino.X {
// 		s.habitantes[idHabitante].posX = int32(s.habitantes[idHabitante].posX) - paso
// 	}

// 	//Movimiento Y
// 	if s.habitantes[idHabitante].posY < destino.Y {
// 		s.habitantes[idHabitante].posY = int32(s.habitantes[idHabitante].posY) + paso
// 	} else if s.habitantes[idHabitante].posY > destino.Y {
// 		s.habitantes[idHabitante].posY = int32(s.habitantes[idHabitante].posY) - paso
// 	}
// }

/*
***	Revisa estados de habitantes y los hace caminar hacia el agua mas cercana
 */
func (s *server) walkToWater(coordenadas []coordenadaAgua) {

	var umbralSed int32 = 30
	var paso int32 = 20 //Si se quiere poder cambiar la velocidad hay que modificar la funcion

	for i := range s.habitantes {

		if s.habitantes[i].estado < umbralSed {

			destino, err := closestWater(s.habitantes[i], coordenadas)
			fmt.Printf("\n\nHabitante %d tiene sed:\n\tEstado: %d\n\tPosicion (%d, %d)\n\tMoviendose a (%d, %d)\n",
				i, s.habitantes[i].estado, s.habitantes[i].posX, s.habitantes[i].posY, destino.X, destino.Y)

			if err != nil {
				log.Fatalf("No se pudo obtener coordenadas de destino para habitante %d\n\tCoordenadas: (%d, %d)\n\tEstado: %d\n\tError: %v",
					i, s.habitantes[i].posX, s.habitantes[i].posY, s.habitantes[i].estado, err)
			}

			//Movimiento X
			if s.habitantes[i].posX < destino.X {
				if s.habitantes[i].posX+paso > destino.X {
					s.habitantes[i].posX = destino.X
				} else {
					s.habitantes[i].posX = int32(s.habitantes[i].posX) + paso
				}
			} else if s.habitantes[i].posX > destino.X {
				if s.habitantes[i].posX-paso < destino.X {
					s.habitantes[i].posX = destino.X
				} else {
					s.habitantes[i].posX = int32(s.habitantes[i].posX) - paso
				}
			}

			//Movimiento Y
			if s.habitantes[i].posY < destino.Y {
				if s.habitantes[i].posY+paso > destino.Y {
					s.habitantes[i].posY = destino.Y
				} else {
					s.habitantes[i].posY = int32(s.habitantes[i].posY) + paso
				}
			} else if s.habitantes[i].posY > destino.Y {
				if s.habitantes[i].posY-paso < destino.Y {
					s.habitantes[i].posY = destino.Y
				} else {
					s.habitantes[i].posY = int32(s.habitantes[i].posY) - paso
				}
			}

			//LLega a coordenada y toma agua
			if s.habitantes[i].posX == destino.X && s.habitantes[i].posY == destino.Y {
				s.habitantes[i].estado += 100
			}
		}
	}

}

// Levantar como servidor mientras pa hacer pruebas (se supone que esto deberia estar en central?)
func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServicioHabitantesServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
