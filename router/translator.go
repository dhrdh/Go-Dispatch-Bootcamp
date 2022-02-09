package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TranslatorController interface {
	Fetch(w http.ResponseWriter, r *http.Request)
	FetchById(w http.ResponseWriter, r *http.Request)
}

func Setup(c TranslatorController) *mux.Router {
	log.Println("In router | Setup")

	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/fetch", c.Fetch).Methods(http.MethodGet).Name("Fetch")
	v1.HandleFunc("/fetch/{id}", c.FetchById).Methods(http.MethodGet).Name("FetchById")

	return r
}
