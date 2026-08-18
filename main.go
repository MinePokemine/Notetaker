package main

import (
	"log"
	"net/http"

	handlers "github.com/MinePokemine/notetaker/Handlers"
	handlers_account "github.com/MinePokemine/notetaker/Handlers/Account"
)

func main() {
	// Web App
	http.HandleFunc("GET /", handlers.LoadHTML("Templates/index.html"))
	http.HandleFunc("GET /account/signup/", handlers.LoadHTML("Templates/signup.html"))
	http.HandleFunc("GET /account/rename/", handlers.LoadHTML("Templates/rename.html"))
	http.HandleFunc("GET /account/", handlers.CreateFuncTemplateHandler("Templates/myaccount.html", handlers_account.MyAccount))

	// API
	http.HandleFunc("POST /api/account/signup/", handlers_account.CreateAccount)
	http.HandleFunc("POST /api/account/rename/", handlers_account.RenameAccount)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
