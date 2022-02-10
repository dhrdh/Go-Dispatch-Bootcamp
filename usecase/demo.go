package usecase

import (
	"Go-Dispatch-Bootcamp/types"
	"errors"
	"fmt"
	"log"
)

type demoService interface {
	GetUsers() (*[]types.User, error)
	GetUsersMap() (map[int]types.User, error)
	GetFeedUsers() ([][]string, error)
	FetchCsvFromRemote() ([][]string, error)
	UpdateUsers(*[]types.User) (bool, error)
}

type demoUsecase struct {
	service demoService
}

func New(s demoService) *demoUsecase {
	log.Println("In usecase | constructor")

	return &demoUsecase{
		service: s,
	}
}

func (tu *demoUsecase) Fetch() (*[]types.User, error) {
	log.Println("In usecase | Fetch")

	users, err := tu.service.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (tu *demoUsecase) FetchById(id int) (*types.User, error) {
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

func (tu *demoUsecase) Feed() ([][]string, error) {
	log.Println("In usecase | Feed")

	return tu.service.GetFeedUsers()
}

func (tu *demoUsecase) UpdateUsersFromFeed() (bool, error) {
	log.Println("In usecase | UpdateUsersFromFeed")

	csvUsers, err := tu.service.FetchCsvFromRemote()

	if err != nil {
		return false, err
	}

	var users []types.User

	for i, csvUser := range csvUsers {
		// skip table title
		if i == 0 {
			continue
		}

		user := types.User{
			Id:         i,
			Username:   csvUser[0],
			Identifier: csvUser[1],
			FirstName:  csvUser[2],
			LastName:   csvUser[3],
		}

		users = append(users, user)
	}

	return tu.service.UpdateUsers(&users)
}
