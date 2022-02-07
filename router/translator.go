package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TranslatorController interface {
	UpdateFromRemote(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
}

func Setup(c TranslatorController) *mux.Router {
	log.Println("In router | Setup")

	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/update-from-remote", c.UpdateFromRemote).Methods(http.MethodGet).Name("UpdateFromRemote")
	v1.HandleFunc("/fetch", c.Fetch).Methods(http.MethodGet).Name("Fetch")

	return r
}
