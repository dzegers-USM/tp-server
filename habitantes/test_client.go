package main

import (
	"context"
	"io"
	"log"

	// "time"

	pb "github.com/dzegers-USM/tp-server/habitantes/misc"

	"google.golang.org/grpc"
)

// Función para mostrar estados iniciales de manera estructurada
func mostrarEstadosIniciales(habitantes []*pb.Habitante) {
	log.Printf("Estados iniciales de los habitantes:\n")
	for i, h := range habitantes {
		log.Printf("Habitante %d: posX = %d, posY = %d, estado = %d", i, h.PosX, h.PosY, h.Estado)
	}
	log.Println()
}

/*
***	Ejemplo con 50 habitantes
 */
func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServicioHabitantesClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*60);
	// defer cancel();
	ctx := context.Background()

	//Pedir a server inicializar estados de habitantes.
	req := &pb.InicializadorRequest{NumHabitantes: 10}
	estadosIniciales, err := c.InicializadorHabitantes(ctx, req)
	if err != nil {
		log.Fatalf("could not initialize estados: %v", err)
	}

	for {
		resp, err := estadosIniciales.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}
		mostrarEstadosIniciales(resp.HabitantesInicial)
	}

	//Inicia y escucha actualizacion de estados (pd: &pb.EstadoRequest no hace nada)
	estadosActuales, err := c.ActualizarEstado(ctx, &pb.EstadoRequest{Request: 50})
	if err != nil {
		log.Fatalf("could not start updates: %v", err)
	}

	for {
		resp, err := estadosActuales.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}
		log.Printf("Estados actuales de los habitantes:\n%v\n\n", resp.GetEstadoHabitante())
	}
}
