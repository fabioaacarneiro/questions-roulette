package handler

import (
	"database/sql"
	"html/template"
	"jogodasperguntas/internal/domain"
	"jogodasperguntas/internal/repository"
	"net/http"
	"strconv"
	"time"
)

type PageData struct {
	Title      string
	Questions  []domain.Question
	ActualYear int
}

type QuestionHandler struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionHandler(db *sql.DB) *QuestionHandler {
	return &QuestionHandler{
		questionRepository: repository.NewQuestionRepository(db),
	}
}

func parseTemplates() (*template.Template, error) {
	return template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/questions.html",
	)
}

func (qh *QuestionHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "parâmetro id não encontrado", http.StatusBadRequest)
		return
	}

	question, err := qh.questionRepository.FindById(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ, err := parseTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Jogo das Perguntas",
		Questions:  []domain.Question{*question},
		ActualYear: time.Now().Year(),
	}

	if err := templ.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (qh *QuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "parâmetro id não encontrado", http.StatusBadRequest)
		return
	}

	if err := qh.questionRepository.Delete(idInt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/questions", http.StatusSeeOther)
}

func (qh *QuestionHandler) GetList(w http.ResponseWriter, r *http.Request) {
	questions, err := qh.questionRepository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ, err := parseTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Jogo das Perguntas",
		Questions:  questions,
		ActualYear: time.Now().Year(),
	}

	if err := templ.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (qh *QuestionHandler) Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	question := domain.Question{
		Question: r.FormValue("question"),
	}

	if err := qh.questionRepository.Store(question); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/questions", http.StatusSeeOther)
}

func (qh *QuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := qh.questionRepository.Update(idInt, r.FormValue("question")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/questions", http.StatusSeeOther)
}
