package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("API Application starting...")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/voorraad/kruiden", handleGetKruiden).Methods("GET")
}

func handleGetKruiden(w http.ResponseWriter, r *http.Request) {
	getKruidenAmount()
}
