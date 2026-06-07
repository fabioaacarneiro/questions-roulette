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

type ListQuestionPageData struct {
	Title      string
	Questions  []domain.Question
	ActualYear int
}

type PageData struct {
	Title      string
	ActualYear int
}

const (
	layoutBase = "web/templates/layouts/base.html"

	partialHeader       = "web/templates/partials/header.html"
	partialFooter       = "web/templates/partials/footer.html"
	partialQuestionItem = "web/templates/partials/question_item.html"
	partialSortResult   = "web/templates/partials/sort_result.html"

	pageQuestions = "web/templates/pages/questions.html"
	pageSort      = "web/templates/pages/sort.html"
)

type QuestionHandler struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionHandler(db *sql.DB) *QuestionHandler {
	return &QuestionHandler{
		questionRepository: repository.NewQuestionRepository(db),
	}
}

func renderQuestionItem(w http.ResponseWriter, question domain.Question) error {
	templ, err := template.ParseFiles(partialQuestionItem)
	if err != nil {
		return err
	}
	return templ.ExecuteTemplate(w, "question-item", question)
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

	templ, err := template.ParseFiles(
		layoutBase,
		partialHeader,
		partialFooter,
		partialQuestionItem,
		pageQuestions,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := ListQuestionPageData{
		Title:      "Jogo das Perguntas",
		Questions:  []domain.Question{*question},
		ActualYear: time.Now().Year(),
	}

	if err := templ.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (qh *QuestionHandler) GetList(w http.ResponseWriter, r *http.Request) {
	questions, err := qh.questionRepository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ, err := template.ParseFiles(
		layoutBase,
		partialHeader,
		partialFooter,
		partialQuestionItem,
		pageQuestions,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := ListQuestionPageData{
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

	question, err := qh.questionRepository.Store(domain.Question{
		Question: r.FormValue("question"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := renderQuestionItem(w, *question); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	question, err := qh.questionRepository.Update(idInt, r.FormValue("question"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := renderQuestionItem(w, *question); err != nil {
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

	w.WriteHeader(http.StatusOK)
}

func (qh *QuestionHandler) Sort(w http.ResponseWriter, r *http.Request) {
	question, err := qh.questionRepository.Sort()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ, err := template.ParseFiles(partialSortResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, question); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (qh *QuestionHandler) SortPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		layoutBase,
		partialHeader,
		partialFooter,
		pageSort,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Jogo das Perguntas",
		ActualYear: time.Now().Year(),
	}

	if err := templ.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
