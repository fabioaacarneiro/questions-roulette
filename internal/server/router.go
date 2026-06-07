package server

import (
	"database/sql"
	"jogodasperguntas/internal/handler"
	"net/http"
)

func SetupRoutes(db *sql.DB) {
	// servindo arquivos da pasta assets
	fs := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// definindo as rotas
	qh := handler.NewQuestionHandler(db)
	http.HandleFunc("POST /question/delete/{id}", qh.Delete)
	http.HandleFunc("POST /question/update/{id}", qh.Update)
	http.HandleFunc("POST /question/store", qh.Store)
	http.HandleFunc("GET /question/{id}", qh.GetById)
	http.HandleFunc("GET /questions", qh.GetList)
}
