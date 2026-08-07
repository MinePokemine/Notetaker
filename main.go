package main

import (
	"log"
	"net/http"

	handlers "github.com/MinePokemine/notetaker/Handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
