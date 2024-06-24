#! /usr/bin/bash
go run init-csv.go
scp coordenadas.csv admin@172.31.6.119:~
go run main.go