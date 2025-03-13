package routes

import (
	"jogodasperguntas/controllers"
	"jogodasperguntas/services"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
	Routes []Route
}

type Route struct {
	Path    string
	Handler func(r *mux.Router)
}

func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

func (r *Router) AddRoute(route Route) {
	r.Routes = append(r.Routes, route)
}

func (r *Router) RegisterRoutes() {
	for _, route := range r.Routes {
		route.Handler(r.Router.PathPrefix(route.Path).Subrouter())
	}
}

func RouterHandler() *mux.Router {
	r := NewRouter()

	assetsDir, _ := filepath.Abs("views/assets")
	fs := http.FileServer(http.Dir(assetsDir))
	r.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	questionController := controllers.NewQuestionsController(services.NewQuestionsService())
	questionRoutes := NewQuestionRoutes(questionController)
	r.AddRoute(Route{Path: "/questions", Handler: questionRoutes.RegisterQuestionRoutes})

	r.RegisterRoutes()

	return r.Router
}
