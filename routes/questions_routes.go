package routes

import (
	"jogodasperguntas/controllers"

	"github.com/gorilla/mux"
)

type QuestionsRoute struct {
	router *controllers.QuestionsController
}

func NewQuestionRoutes(controller *controllers.QuestionsController) *QuestionsRoute {
	return &QuestionsRoute{router: controller}
}

func (c *QuestionsRoute) RegisterQuestionRoutes(r *mux.Router) {
	r.HandleFunc("", c.router.GetAllQuestions).Methods("GET")
	r.HandleFunc("/get/{id}", c.router.GetQuestionById).Methods("GET")
	r.HandleFunc("/update/{id}", c.router.UpdateQuestion).Methods("POST")
	r.HandleFunc("/delete/{id}", c.router.DeleteQuestion).Methods("POST")
	r.HandleFunc("/store", c.router.StoreQuestion).Methods("POST")
}
