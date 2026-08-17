package main

import (
	"log"
	"net/http"

	handlers "github.com/MinePokemine/notetaker/Handlers"
	handlers_account "github.com/MinePokemine/notetaker/Handlers/Account"
	helpers "github.com/MinePokemine/notetaker/Helpers"
)

func main() {
	http.HandleFunc("GET /", handlers.CreateTemplateHandler("/Templates/index.html", helpers.Null{}))
	http.HandleFunc("POST /api/account/signup/", handlers_account.CreateAccount)
	http.HandleFunc("POST /api/account/rename/", handlers_account.RenameAccount)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
