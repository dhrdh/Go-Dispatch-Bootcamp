package service

import (
	"Go-Dispatch-Bootcamp/types"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const sampleCsvUrl = "https://support.staffbase.com/hc/en-us/article_attachments/360009197031/username.csv"
const fileName = "data.csv"

type translatorService struct {}

func New() *translatorService {
	log.Println("In service | constructor")

	return &translatorService{}
}

func (ts *translatorService) FetchCsvFromRemote() (*types.Csv, error) {
	url := sampleCsvUrl

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := types.Csv {
		Lines: data,
	}

	return &result, nil
}

func (ts *translatorService) SaveCsvToFile(doc *types.Csv) (bool, error) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return false, err
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	err = w.WriteAll(doc.Lines)
	if err != nil {
		return false, err
	}

	log.Println("data.csv was updated")

	return true, nil
}

func (ts *translatorService) ReadCsvFromFile() (*types.Csv, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &types.Csv{ Lines: records }, nil
}

type user struct {
	Username string `json:"username"`
	Identifier string `json:"identifier"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func (ts *translatorService) ConvertCsvToJson(doc *types.Csv) (*types.Json, error) {
	var users []user

	for i, line := range doc.Lines {
		// skip table title
		if i == 0 {
			continue
		}

		users = append(users, user {
			line[0],
			line[1],
			line[2],
			line[3],
		})
	}

	data, err := json.Marshal(users)

	if err != nil {
		return nil, err
	}

	result := types.Json {
		Data: data,
	}
	return &result, nil
}
