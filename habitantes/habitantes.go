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

	pb "habitantes/misc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServicioHabitantesServer
	estados []int32
	mu      sync.Mutex // Protege el acceso concurrente a las variables
}

/*
**	Set niveles de sed iniciales para cada habitante y comunicarlos a cliente.
**	Cliente debe comunicar el numero de habitantes a setear en su request.
 */
func (s *server) InicializadorHabitantes(in *pb.InicializadorRequest, stream pb.ServicioHabitantes_InicializadorHabitantesServer) error {

	//Rango de sed inicial
	const min int32 = 5
	const max int32 = 100

	//Protege el acceso concurrente a las variables mientras se inicializan.
	s.mu.Lock()
	defer s.mu.Unlock()

	//Set niveles de sed iniciales para cada habitante y comunicarlos a cliente.
	s.estados = make([]int32, in.GetNumHabitantes())
	fmt.Printf("Iniciando con n=%d habitantes", in.GetNumHabitantes())
	for i := range s.estados {
		s.estados[i] = rand.Int31n(max-min+1) + min
	}

	//Enviar stream de estado inicial a cliente
	response := &pb.InicializadorResponse{EstadosIniciales: s.estados}
	if err := stream.Send(response); err != nil {
		return err
	}

	return nil
}

/*
***	Cada 5 segundos actualiza el estado de los habitantes (resta 1 a cada uno) y se lo informa al cliente.
 */
func (s *server) ActualizarEstado(in *pb.EstadoRequest, stream pb.ServicioHabitantes_ActualizarEstadoServer) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			s.updateEstados()
			response := &pb.EstadoResponse{Respuesta: s.estados}
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
	for i := range s.estados {
		if s.estados[i] > 0 {
			s.estados[i] -= 1
		}
	}
}

/*
***	Cada 5 segundos actualiza el estado de los habitantes (resta 1 a cada uno) y se lo informa al cliente.
 */
func (s *server) ConsumirRecurso(in *pb.EstadoRequest, stream pb.ServicioHabitantes_ConsumirRecursoServer) error {
	//Protege el acceso concurrente a las variables mientras se ejecuta la funcion.
	s.mu.Lock()
	defer s.mu.Unlock()

	var aumento int32 = in.GetRequest()

	for i := range s.estados {
		if s.estados[i] > 0 {
			s.estados[i] += aumento
		}
	}

	response := &pb.EstadoResponse{Respuesta: s.estados}
	if err := stream.Send(response); err != nil {
		return err
	}
	return stream.Context().Err()
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
