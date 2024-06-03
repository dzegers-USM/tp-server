package main

import (
	// "context"
	"fmt"
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

type habitante struct {
	posX   int32
	posY   int32
	estado int32
}

type server struct {
	pb.UnimplementedServicioHabitantesServer
	habitantes []habitante
	mu         sync.Mutex // Protege el acceso concurrente a las variables
}

/*
**	Set parametros iniciales para cada habitante y comunicarlos a cliente.
**	Cliente debe comunicar el numero de habitantes a setear en su request.
 */
func (s *server) InicializadorHabitantes(in *pb.InicializadorRequest, stream pb.ServicioHabitantes_InicializadorHabitantesServer) error {

	//Parametros
	const min_posX int32 = 0
	const max_posX int32 = 600
	const min_posY int32 = 0
	const max_posY int32 = 600
	const min_sed int32 = 0
	const max_sed int32 = 100

	//Protege el acceso concurrente a las variables mientras se inicializan.
	s.mu.Lock()
	defer s.mu.Unlock()

	//Set parametros iniciales para cada habitante y comunicarlos a cliente.
	s.habitantes = make([]habitante, in.GetNumHabitantes())
	for i := range s.habitantes {
		s.habitantes[i] = habitante{
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
		}
	}
	return responseHabitantes
}

/*
***	Mueve a los habitantes de manera aleatorea cada 1 seg.
***	Cada 5 segundos actualiza el nivel de sed de los habitantes (resta 1 a cada uno) y se lo informa al cliente.
 */
func (s *server) ActualizarEstado(in *pb.EstadoRequest, stream pb.ServicioHabitantes_ActualizarEstadoServer) error {
	ticker_sed := time.NewTicker(5000 * time.Millisecond) //n(000) Segundos
	defer ticker_sed.Stop()

	ticker_movimiento := time.NewTicker(500 * time.Millisecond) //n(000) Segundos
	defer ticker_movimiento.Stop()

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
***	Disminuye estado de habitantes en -1.
 */
func (s *server) updateEstados() {
	for i := range s.habitantes {
		if s.habitantes[i].estado > 0 {
			s.habitantes[i].estado -= 1
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
	for i := range s.habitantes {
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
