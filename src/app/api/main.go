package main

import (
	"fmt"
	"log"

	"padaria/src/app/api/config"
	"padaria/src/app/api/endpoints/router"
	"padaria/src/infra/postgres"

	_ "padaria/src/app/api/docs"
)

// @title Padaria API
// @version 1.0
// @description This is an example bakery server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost
// @BasePath /api
func main() {
	setupPostgres()
	serveAddress(config.ServerHost, config.ServerPort)
}

func serveAddress(host string, port int) {
	server := router.Start()

	address := fmt.Sprintf("%s:%d", host, port)

	server.Logger.Fatal(server.Start(address))
}

func setupPostgres() {
	err := postgres.SetUpCredentials(
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDBName,
		config.PostgresHost,
		config.PostgresPort,
	)
	if err != nil {
		log.Fatal(err)
	}
}
