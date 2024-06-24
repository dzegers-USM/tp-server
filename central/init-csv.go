package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"

	pb "github.com/dzegers-USM/tp-server/agua/server"
	"google.golang.org/grpc"
)

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

func getCoordenadasWater(client pb.MiServicioClient) {
	resp, err := client.EnviarCoordenadas(context.Background(), &pb.CoordenadasRequest{})
	if err != nil {
		log.Fatalf("Failed to invoke EnviarCoordenadas from water sv: %v", err)
	}
	log.Printf("Respuesta recibida:\n%s", resp.Respuesta)

	// Ruta del directorio donde se guardará el archivo CSV
	directory := `./` // Cambia esta ruta por la que necesites

	// Nombre del archivo CSV de salida en el directorio específico
	filename := filepath.Join(directory, "coordenadas.csv")

	// Escribir datos en el archivo CSV
	errcsv := writeCSV(resp.Respuesta, filename)
	if errcsv != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %v", errcsv)
	}

	log.Printf("Datos escritos en el archivo %s exitosamente.", filename)
}

func main() {
	var ip_agua = "172.31.1.228"
	conn, err := grpc.Dial(ip_agua+":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMiServicioClient(conn)

	getCoordenadasWater(client)
}
