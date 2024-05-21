package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	pb "github.com/dzegers-USM/tp-server/agua/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initWater(client pb.MiServicioClient, n int32) {
	client.Inicializador(context.Background(), &pb.InicializadorRequest{Inicializador: n})
	log.Printf("Init water with %d tiles\n", n)
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
		log.Printf("Got msg from water sv: %v\n", msg.Estado)
	}
}

func main() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	type SimulationConfig struct {
		WaterTiles int
	}
	var config SimulationConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
	log.Printf("Starting with config: %+v", config)

	waterConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to water resource sv: %v", err)
	}
	defer waterConn.Close()
	waterClient := pb.NewMiServicioClient(waterConn)

	// Enviar parámetros iniciales
	initWater(waterClient, int32(config.WaterTiles))

	// Monitorear heartbeat de los sv de recursos
	watchWater(waterClient)
}
