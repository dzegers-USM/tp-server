package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dzegers-USM/tp-server/agua/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// Importa el paquete generado por protoc
)

func initWater(client pb.MiServicioClient, n int32) {
	client.Inicializador(context.Background(), &pb.InicializadorRequest{Inicializador: n})
	fmt.Printf("Init water with n=%d\n", n)
}

func watchWater(client pb.MiServicioClient) {
	stream, err := client.Heartbeat(context.Background(), &pb.HeartbeatRequest{})
	if err != nil {
		log.Fatalf("Failed to invoke heartbeat from water sv: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error watching water sv: %v", err)
		}
		fmt.Printf("Got msg from water sv: %v\n", msg.Estado)
	}
}

func main() {
	var nWater int
	fmt.Print("Water resource availability (int): ")
	fmt.Scan(&nWater)

	waterConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to water resource sv: %v", err)
	}
	defer waterConn.Close()
	waterClient := pb.NewMiServicioClient(waterConn)

	// Enviar parámetros iniciales
	initWater(waterClient, int32(nWater))

	// Monitorear heartbeat de los sv de recursos
	watchWater(waterClient)
}
