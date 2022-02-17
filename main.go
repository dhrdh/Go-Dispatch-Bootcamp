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
	demoService := service.New()
	demoUsecase := usecase.New(demoService)
	demoController := controller.New(demoUsecase)
	httpRouter := router.Setup(demoController)

	err := http.ListenAndServe("localhost:8080", httpRouter)

	if err != nil {
		log.Printf("Server start error: %v", err)
	}

	log.Println("Server is started on port 8080")
}
