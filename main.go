package main

import (
	"fmt"
	"jogodasperguntas/config"
	"jogodasperguntas/routes"
	"log"
	"net/http"
)

func main() {
	config.Connect()

	router := routes.RouterHandler()

	port := ":8080"
	fmt.Println("Servidor rodando em http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
