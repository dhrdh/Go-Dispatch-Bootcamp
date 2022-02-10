package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"Go-Dispatch-Bootcamp/types"
)

type usecase interface {
	Fetch() (*[]types.User, error)
	FetchById(int) (*types.User, error)
	Feed() (*[]types.FeedUser, error)
}

type translatorController struct {
	usecase usecase
}

func (tc *translatorController) Fetch(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | Fetch")

	users, err := tc.usecase.Fetch()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Fetch error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(users)
	w.Write(data)
}

func (tc *translatorController) FetchById(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | FetchById")

	vars := mux.Vars(r)
	stringId, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "FetchById error: id is not defined")
		return
	}

	id, err := strconv.Atoi(stringId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("FetchById error: Id '%v' is not a number", stringId))
		return
	}

	user, err := tc.usecase.FetchById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("FetchById error: %v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(user)
	w.Write(data)
}

func (tc *translatorController) Feed(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | Feed")

	users, err := tc.usecase.Feed()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Feed error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(users)
	w.Write(data)
}

func New(uc usecase) *translatorController {
	log.Println("In controller | constructor")

	return &translatorController{
		usecase: uc,
	}
}
