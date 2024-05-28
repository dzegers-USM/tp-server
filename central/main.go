package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	pb "github.com/dzegers-USM/tp-server/agua/server"
	pbh "github.com/dzegers-USM/tp-server/habitantes/misc"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initHabitantes(client pbh.ServicioHabitantesClient, q amqp.Queue, ch *amqp.Channel) {
	estadosIniciales, err := client.InicializadorHabitantes(context.Background(), &pbh.InicializadorRequest{NumHabitantes: 50})
	if err != nil {
		log.Fatalf("Could not initialize estados: %v", err)
	}

	for {
		resp, err := estadosIniciales.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}

		data, _ := json.Marshal(resp.EstadosIniciales)

		log.Printf("Estados iniciales de los habitantes:\n%v\n\n", resp.EstadosIniciales)
		err = ch.PublishWithContext(context.Background(),
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)
		if err != nil {
			log.Fatalf("Failed to publish a message: %v", err)
		}
	}
}

func watchHabitantes(client pbh.ServicioHabitantesClient, q amqp.Queue, ch *amqp.Channel) {
	estadosActuales, err := client.ActualizarEstado(context.Background(), &pbh.EstadoRequest{Request: 50})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	for {
		resp, err := estadosActuales.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}

		data, _ := json.Marshal(resp.Respuesta)

		log.Printf("Estados actuales de los habitantes:\n%v\n\n", resp.Respuesta)
		err = ch.PublishWithContext(context.Background(),
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)
		if err != nil {
			log.Fatalf("Failed to publish a message: %v", err)
		}
	}
}

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

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	amqpCh, err := amqpConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a Rabbit channel: %v", err)
	}

	q, err := amqpCh.QueueDeclare(
		"Habitantes", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)

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
	go watchWater(waterClient)

	habitantesConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer habitantesConn.Close()
	habitantesClient := pbh.NewServicioHabitantesClient(habitantesConn)

	initHabitantes(habitantesClient, q, amqpCh)

	watchHabitantes(habitantesClient, q, amqpCh)
}
