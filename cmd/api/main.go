package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wexel-nath/pets/pkg/api"
	"github.com/wexel-nath/pets/pkg/config"
	"github.com/wexel-nath/pets/pkg/database"
)

func main() {
	config.Configure()
	database.GetConnection()

	startServer()
}

func startServer() {
	address := ":8080"
	fmt.Println("Listening on " + address)
	log.Fatal(http.ListenAndServe(address, api.GetRouter()))
}
