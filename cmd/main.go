package main

import (
	"log"

	"github.com/trmyxx/AuthService/initializers"
	"github.com/trmyxx/AuthService/internal/router"
	"github.com/trmyxx/AuthService/internal/service"
	"github.com/trmyxx/AuthService/internal/storage"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDataBase()
}

func main() {

	storage := storage.NewStorage()
	service := service.NewService(*storage)

	rest := router.NewRouter(*storage, *service)

	err := rest.Run()
	if err != nil {
		log.Fatal("failed run server", err)
	}
}
