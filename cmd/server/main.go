package main

import (
	"fmt"
	"jogodasperguntas/internal/database"
	"jogodasperguntas/internal/server"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server.SetupRoutes(db)

	port := ":8080"
	fmt.Println("Servidor rodando em http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
