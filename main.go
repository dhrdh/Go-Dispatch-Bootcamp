package main

import (
	"log"
	"net/http"

	"Go-Dispatch-Bootcamp/controller"
	"Go-Dispatch-Bootcamp/router"
	"Go-Dispatch-Bootcamp/service"
	"Go-Dispatch-Bootcamp/usecase"
)

func main() {
	translatorService := service.New()
	translatorUsecase := usecase.New(translatorService)
	translatorController := controller.New(translatorUsecase)
	httpRouter := router.Setup(translatorController)

	err := http.ListenAndServe("localhost:8080", httpRouter)

	if err != nil {
		log.Printf("Server start error: %v", err)
	}

	log.Println("Server is started on port 8080")
}
