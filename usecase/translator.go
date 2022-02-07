package usecase

import (
	"fmt"
	"log"
	"Go-Dispatch-Bootcamp/types"
)

type translatorService interface {
	FetchCsvFromRemote() (*types.Csv, error)
	SaveCsvToFile(*types.Csv) (bool, error)
	ReadCsvFromFile() (*types.Csv, error)
	ConvertCsvToJson(doc *types.Csv) (*types.Json, error)
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

func (tu *translatorUsecase) FetchCsvFromRemote() (bool, error) {
	log.Println("In usecase | FetchCsvFromRemote")

	csv, err := tu.service.FetchCsvFromRemote()

	if err != nil {
		return false, fmt.Errorf("FetchCsvFromRemote error: %v", err)
	}

	success, err := tu.service.SaveCsvToFile(csv)

	if err != nil {
		return false, fmt.Errorf("FetchCsvFromRemote error: %v", err)
	}

	return success, nil
}

func (tu *translatorUsecase) FetchJson() (*types.Json, error) {
	log.Println("In usecase | FetchCsvFromRemote")

	csvDoc, err := tu.service.ReadCsvFromFile()

	if err != nil {
		return nil, fmt.Errorf("FetchCsvFromRemote error: %v", err)
	}

	jsonDoc, err := tu.service.ConvertCsvToJson(csvDoc)

	if err != nil {
		return nil, fmt.Errorf("FetchCsvFromRemote error: %v", err)
	}

	return jsonDoc, nil
}
