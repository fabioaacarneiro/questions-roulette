package controllers

import (
	"jogodasperguntas/dto"
	"jogodasperguntas/services"
	"jogodasperguntas/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type QuestionsController struct {
	service *services.QuestionsService
}

func NewQuestionsController(service *services.QuestionsService) *QuestionsController {
	return &QuestionsController{service: service}
}

func (c *QuestionsController) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := c.service.GetAllQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Jogo das Perguntas",
		"Questions": questions,
	}

	tmplFiles := []string{
		"views/layouts/base.html",
		"views/partials/header.html",
		"views/questions.html",
		"views/partials/footer.html",
	}

	utils.RenderTemplate(w, tmplFiles, data)
}

func (c *QuestionsController) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "parâmetro id não encontrado", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	question := r.FormValue("question")

	err := c.service.UpdateQuestion(id, question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/questions", http.StatusSeeOther)
}

func (c *QuestionsController) GetQuestionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "parâmetro id não encontrado", http.StatusBadRequest)
		return
	}

	question, err := c.service.FindQuestionById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	questions := []dto.Question{}
	questions = append(questions, *question)

	data := map[string]interface{}{
		"Title":     "Jogo das Perguntas",
		"Questions": questions,
	}

	tmplFiles := []string{
		"views/layouts/base.html",
		"views/partials/header.html",
		"views/questions.html",
		"views/partials/footer.html",
	}

	utils.RenderTemplate(w, tmplFiles, data)
}

func (c *QuestionsController) StoreQuestion(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questionDTO := dto.Question{
		Question: r.FormValue("question"),
	}

	if err := c.service.StoreQuestion(questionDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/questions", http.StatusSeeOther)
}

func (c *QuestionsController) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "parâmetro id não encontrado", http.StatusBadRequest)
		return
	}

	err := c.service.DeleteQuestion(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/questions", http.StatusSeeOther)
}
