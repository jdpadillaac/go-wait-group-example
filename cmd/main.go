package main

import (
	appmysql "github.com/jdpadillaac/go-waitgroups-example/infrastructure/mysql"
	restapi "github.com/jdpadillaac/go-waitgroups-example/infrastructure/rest_api"
	"github.com/jdpadillaac/go-waitgroups-example/internal"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file doest not exist")
	}

	config := internal.SetAppConfig()
	appmysql.SetCnn(config)
	restapi.Start(config)
}
