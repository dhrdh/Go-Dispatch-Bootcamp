package controller

import (
	"fmt"
	"log"
	"net/http"

	"Go-Dispatch-Bootcamp/types"
)

type usecase interface {
	FetchCsvFromRemote() (bool, error)
	FetchJson() (*types.Json, error)
}

type translatorController struct {
	usecase usecase
}

func (tc *translatorController) UpdateFromRemote(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | UpdateFromRemote")

	success, err := tc.usecase.FetchCsvFromRemote()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)

		log.Fatalf("FetchCsvFromRemote error: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{ success: { %v } }", success)))
}

func (tc *translatorController) Fetch(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | Fetch")

	result, err := tc.usecase.FetchJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)

		log.Fatalf("FetchCsvFromRemote error: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result.Data)
}

func New(uc usecase) *translatorController {
	log.Println("In controller | constructor")

	return &translatorController{
		usecase: uc,
	}
}
