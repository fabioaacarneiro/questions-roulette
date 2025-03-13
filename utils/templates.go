package utils

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmplFiles []string, data interface{}) {

	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println("erro ao executar template: ", err)
		http.Error(w, "erro ao renderizar templates", http.StatusInternalServerError)
	}
}
