package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"Go-Dispatch-Bootcamp/types"
)

type demoService struct{}

func New() *demoService {
	return &demoService{}
}

func (ts *demoService) readCsvFromFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("can not open file")
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("read file error")
	}
	return records, nil
}

func (ts *demoService) FetchCsvFromRemote(feedUrl string) ([][]string, error) {
	url := feedUrl

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("can not read csv from remote")
	}

	return data, nil
}

func (ts *demoService) UpdateUsers(users *[]types.User, dataFileName string) (bool, error) {
	file, err := os.Create(dataFileName)
	defer file.Close()
	if err != nil {
		return false, errors.New(fmt.Sprintf("can not open file: %v", dataFileName))
	}

	var csvUsers [][]string

	for _, user := range *users {
		csvUser := []string{
			strconv.Itoa(user.Id),
			user.Username,
			user.Identifier,
			user.FirstName,
			user.LastName,
		}

		csvUsers = append(csvUsers, csvUser)
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	err = w.WriteAll(csvUsers)
	if err != nil {
		return false, err
	}

	log.Println(fmt.Sprintf("%v was updated", dataFileName))

	return true, nil
}

func (ts *demoService) GetUsers(dataFileName string) (*[]types.User, error) {
	records, err := ts.readCsvFromFile(dataFileName)
	if err != nil {
		return nil, err
	}

	var users []types.User

	for _, line := range records {
		id, err := strconv.Atoi(line[0])

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Id '%v' is not a number", line[0]))
		}

		users = append(users, types.User{
			Id:         id,
			Username:   line[1],
			Identifier: line[2],
			FirstName:  line[3],
			LastName:   line[4],
		})
	}

	return &users, nil
}

func (ts *demoService) GetUsersMap(dataFileName string) (map[int]types.User, error) {
	records, err := ts.readCsvFromFile(dataFileName)
	if err != nil {
		return nil, err
	}

	users := make(map[int]types.User, len(records))

	for _, line := range records {
		id, err := strconv.Atoi(line[0])

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Id '%v' is not a number", line[0]))
		}

		users[id] = types.User{
			Id:         id,
			Username:   line[1],
			Identifier: line[2],
			FirstName:  line[3],
			LastName:   line[4],
		}
	}

	return users, nil
}

func (ts *demoService) GetFeedUsers(feedFileName string) ([][]string, error) {
	return ts.readCsvFromFile(feedFileName)
}
