package usecase

import (
	"Go-Dispatch-Bootcamp/types"
	"errors"
	"fmt"
	"log"
)

type translatorService interface {
	GetUsers() (*[]types.User, error)
	GetUsersMap() (map[int]types.User, error)
	GetFeedUsers() (*[]types.FeedUser, error)
}

type translatorUsecase struct {
	service translatorService
}

func New(s translatorService) *translatorUsecase {
	log.Println("In usecase | constructor")

	return &translatorUsecase{
		service: s,
	}
}

func (tu *translatorUsecase) Fetch() (*[]types.User, error) {
	log.Println("In usecase | Fetch")

	users, err := tu.service.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (tu *translatorUsecase) FetchById(id int) (*types.User, error) {
	log.Println("In usecase | FetchById")

	users, err := tu.service.GetUsersMap()

	if err != nil {
		return nil, err
	}

	result, ok := users[id]

	if !ok {
		return nil, errors.New(fmt.Sprintf("User with id: %v was not found", id))
	}

	return &result, nil
}

func (tu *translatorUsecase) Feed() (*[]types.FeedUser, error) {
	log.Println("In usecase | Fetch")

	users, err := tu.service.GetFeedUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
