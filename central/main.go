package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	pb "github.com/dzegers-USM/tp-server/agua/server"
	pbh "github.com/dzegers-USM/tp-server/habitantes/misc"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Función para mostrar estados de manera estructurada
func mostrarEstados(habitantes []*pbh.Habitante) {
	log.Printf("Estados de los habitantes:\n")
	for i, h := range habitantes {
		log.Printf("Habitante %d: posX = %d, posY = %d, estado = %d, id = %d", i, h.PosX, h.PosY, h.Estado, h.Id)
	}
	log.Println()
}

func initHabitantes(client pbh.ServicioHabitantesClient, q amqp.Queue, ch *amqp.Channel) {
	estadosIniciales, err := client.InicializadorHabitantes(context.Background(), &pbh.InicializadorRequest{NumHabitantes: 20})
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
		mostrarEstados(resp.HabitantesInicial)

		data, _ := json.Marshal(resp.HabitantesInicial)
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
		mostrarEstados(resp.GetEstadoHabitante())

		data, _ := json.Marshal(resp.GetEstadoHabitante())
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

func getCoordenadasWater(client pb.MiServicioClient) {
	resp, err := client.EnviarCoordenadas(context.Background(), &pb.CoordenadasRequest{})
	if err != nil {
		log.Fatalf("Failed to invoke EnviarCoordenadas from water sv: %v", err)
	}
	log.Printf("Respuesta recibida:\n%s", resp.Respuesta)

	// Ruta del directorio donde se guardará el archivo CSV
	directory := `../../Civitas` // Cambia esta ruta por la que necesites

	// Nombre del archivo CSV de salida en el directorio específico
	filename := filepath.Join(directory, "coordenadas.csv")

	// Escribir datos en el archivo CSV
	errcsv := writeCSV(resp.Respuesta, filename)
	if errcsv != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %v", errcsv)
	}

	log.Printf("Datos escritos en el archivo %s exitosamente.", filename)
}

func writeCSV(data string, filename string) error {
	// Abre el archivo CSV para escritura
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crea un escritor CSV
	writer := csv.NewWriter(file)
	writer.Comma = ';'

	//print(data, "\n")

	// Divide el mensaje en líneas
	lines := strings.Split(data, "\n")

	//print(lines, "\n")

	// Escribe cada línea en el archivo CSV
	for _, line := range lines {
		// Divide cada línea en campos usando ";"
		fields := strings.Split(line, ";")
		print("xd1", fields, "\n")

		if len(fields) == 1 {
			continue
		}
		// Escribe los campos en el archivo CSV
		err := writer.Write(fields)
		if err != nil {
			return err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}

	return nil
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMiServicioClient(conn)

	getCoordenadasWater(client)

	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	amqpConn, err := amqp.Dial("amqps://aoqlvbcv:cIAh4WRAfs13b8N_Ooq7YMuC5IoUzHLd@prawn.rmq.cloudamqp.com/aoqlvbcv")
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
