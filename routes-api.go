package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	//r.Route("/api", func(mux chi.Router) {
	//	// add any api routes here
	//	r.Get("/test-api", func(writer http.ResponseWriter, request *http.Request) {
	//		var payload struct {
	//			Content string `json:"content"`
	//		}
	//		payload.Content = "Hello from API Routes"
	//		a.App.WriteJSON(writer, http.StatusOK, payload)
	//	})
	//})
	return r
}
