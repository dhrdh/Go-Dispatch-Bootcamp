package usecase

import (
	"errors"
	"fmt"

	"Go-Dispatch-Bootcamp/types"
)

const feedUrl = "http://localhost:8080/api/v1/feed"
const dataFileName = "data/data.csv"
const feedFileName = "data/feed.csv"

type demoService interface {
	GetUsers(string) (*[]types.User, error)
	GetUsersConcurrently(string, string, int, int) (*[]types.User, error)
	GetUsersMap(string) (map[int]types.User, error)
	GetFeedUsers(string) ([][]string, error)
	FetchCsvFromRemote(string) ([][]string, error)
	UpdateUsers(*[]types.User, string) (bool, error)
}

type demoUsecase struct {
	service demoService
}

// New - creates new instance of the usecase.
func New(s demoService) *demoUsecase {
	return &demoUsecase{
		service: s,
	}
}

// Fetch - returns all users.
func (tu *demoUsecase) Fetch() (*[]types.User, error) {
	users, err := tu.service.GetUsers(dataFileName)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// FetchConcurrently - returns required amount of users in concurrent mode.
func (tu *demoUsecase) FetchConcurrently(idType string, items int, itemsPerWorker int) (*[]types.User, error) {
	users, err := tu.service.GetUsersConcurrently(dataFileName, idType, items, itemsPerWorker)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// FetchById - returns user by id.
func (tu *demoUsecase) FetchById(id int) (*types.User, error) {
	users, err := tu.service.GetUsersMap(dataFileName)

	if err != nil {
		return nil, err
	}

	result, ok := users[id]

	if !ok {
		return nil, errors.New(fmt.Sprintf("User with id: %v was not found", id))
	}

	return &result, nil
}

// Feed - returns data in csv format. Contains feed for users.
func (tu *demoUsecase) Feed() ([][]string, error) {
	return tu.service.GetFeedUsers(feedFileName)
}

// UpdateUsersFromFeed - makes call to feed endpoint and update users in data.csv.
func (tu *demoUsecase) UpdateUsersFromFeed() (bool, error) {
	csvUsers, err := tu.service.FetchCsvFromRemote(feedUrl)

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

	return tu.service.UpdateUsers(&users, dataFileName)
}
