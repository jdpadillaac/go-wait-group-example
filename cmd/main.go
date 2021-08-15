package main

import (
	"fmt"
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

	fmt.Println(config.Version)

}
