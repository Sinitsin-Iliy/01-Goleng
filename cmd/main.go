package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/zakonm/pkg/database"
	"example.com/zakonm/pkg/routes"
)

func main() {
	database.ConnectToDatabase()

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	fmt.Println("Сервер запущен на порту :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
