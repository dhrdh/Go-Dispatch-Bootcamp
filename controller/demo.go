package controller

import (
	"bytes"
	"encoding/csv"
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
	Feed() ([][]string, error)
	UpdateUsersFromFeed() (bool, error)
}

type demoController struct {
	usecase usecase
}

func (tc *demoController) Fetch(w http.ResponseWriter, r *http.Request) {
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

func (tc *demoController) FetchById(w http.ResponseWriter, r *http.Request) {
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

func (tc *demoController) Feed(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | Feed")

	users, err := tc.usecase.Feed()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Feed error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	b := new(bytes.Buffer)
	csvWriter := csv.NewWriter(b)
	csvWriter.WriteAll(users)
	w.Write(b.Bytes())
}

func (tc *demoController) UpdateUsersFromFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller | UpdateUsersFromFeed")

	success, err := tc.usecase.UpdateUsersFromFeed()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "UpdateUsersFromFeed error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{ success: { %v } }", success)))
}

func New(uc usecase) *demoController {
	log.Println("In controller | constructor")

	return &demoController{
		usecase: uc,
	}
}
