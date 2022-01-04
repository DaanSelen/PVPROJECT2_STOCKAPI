package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initSys()
	log.Println("API Application starting...")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/voorraad/all/brood", handleGetAllBrood).Methods("GET")
	myRouter.HandleFunc("/voorraad/brood", handleGetBroodWQuery).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/broodbeleg", handleGetAllBroodbeleg).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/fruit", handleGetAllFruit).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/kruid", handleGetAllKruid).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/snoep", handleGetAllSnoep).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/vlees", handleGetAllVlees).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/zuivel", handleGetAllZuivel).Methods("GET")

	http.ListenAndServe(":61909", myRouter)
}

func handleGetAllBrood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedBroden := getAllProducts("brood")
	json.NewEncoder(w).Encode(retrievedBroden)
}

func handleGetAllBroodbeleg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruiden := getAllProducts("broodbeleg")
	json.NewEncoder(w).Encode(retrievedKruiden)
}

func handleGetAllFruit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruiden := getAllProducts("fruit")
	json.NewEncoder(w).Encode(retrievedKruiden)
}

func handleGetAllKruid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruiden := getAllProducts("kruid")
	json.NewEncoder(w).Encode(retrievedKruiden)
}

func handleGetAllSnoep(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruiden := getAllProducts("snoep")
	json.NewEncoder(w).Encode(retrievedKruiden)
}

func handleGetAllVlees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruiden := getAllProducts("vlees")
	json.NewEncoder(w).Encode(retrievedKruiden)
}

func handleGetAllZuivel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruiden := getAllProducts("zuivel")
	json.NewEncoder(w).Encode(retrievedKruiden)
}
