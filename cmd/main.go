package main

import (
	"Auth/internal/router"
	"Auth/internal/service"
	"Auth/internal/storage"
	"log"
)

func main() {

	storage := storage.NewStorage()
	service := service.NewService()

	rest := router.NewRouter(*storage, *service)

	err := rest.Run()
	if err != nil {
		log.Fatal("failed run server", err)
	}
}
