package main

import (
	"log"
	"todos/cli"
	"todos/services"
)
func main() {
	storageService := services.NewJSONStorage("./storage.json")
	todoService, err := services.NewTodoService(*storageService)
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.Start(todoService); err != nil {
		log.Fatal("error while starting cli app : " + err.Error())
	}
}