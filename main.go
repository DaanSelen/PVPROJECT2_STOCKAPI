package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	cat1 = "brood"
	cat2 = "broodbeleg"
	cat3 = "fruit"
	cat4 = "kruid"
	cat5 = "snoep"
	cat6 = "vlees"
	cat7 = "zuivel"
)

func main() {
	initSys()
	log.Println("API Application starting...")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/voorraad", handleVoorraadRoot).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/brood", handleGetAllBrood).Methods("GET")
	myRouter.HandleFunc("/voorraad/brood", handleGetBroodWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/brood", handlePostBroodWQuery).Methods("POST")

	myRouter.HandleFunc("/voorraad/all/broodbeleg", handleGetAllBroodbeleg).Methods("GET")
	myRouter.HandleFunc("/voorraad/broodbeleg", handleGetBroodbelegWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/broodbeleg", handlePostBroodbelegWQuery).Methods("POST")

	myRouter.HandleFunc("/voorraad/all/fruit", handleGetAllFruit).Methods("GET")
	myRouter.HandleFunc("/voorraad/fruit", handleGetFruitWQuery).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/kruid", handleGetAllKruid).Methods("GET")
	myRouter.HandleFunc("/voorraad/kruid", handleGetKruidWQuery).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/snoep", handleGetAllSnoep).Methods("GET")
	myRouter.HandleFunc("/voorraad/snoep", handleGetSnoepWQuery).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/vlees", handleGetAllVlees).Methods("GET")
	myRouter.HandleFunc("/voorraad/vlees", handleGetVleesWQuery).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/zuivel", handleGetAllZuivel).Methods("GET")
	myRouter.HandleFunc("/voorraad/zuivel", handleGetZuivelWQuery).Methods("GET")

	http.ListenAndServe(":61909", myRouter)
}

//ROOT RESPONSE
func handleVoorraadRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Root directory for the voorraad tree.")
}

//ALL PRODUCTS WITHOUT QUERY
func handleGetAllBrood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedBroden := getAllProducts(cat1)
	json.NewEncoder(w).Encode(retrievedBroden)
}
func handleGetAllBroodbeleg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedBroodbeleg := getAllProducts(cat2)
	json.NewEncoder(w).Encode(retrievedBroodbeleg)
}
func handleGetAllFruit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedFruit := getAllProducts(cat3)
	json.NewEncoder(w).Encode(retrievedFruit)
}
func handleGetAllKruid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedKruid := getAllProducts(cat4)
	json.NewEncoder(w).Encode(retrievedKruid)
}
func handleGetAllSnoep(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedSnoep := getAllProducts(cat5)
	json.NewEncoder(w).Encode(retrievedSnoep)
}
func handleGetAllVlees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedVlees := getAllProducts(cat6)
	json.NewEncoder(w).Encode(retrievedVlees)
}
func handleGetAllZuivel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	retrievedZuivel := getAllProducts(cat7)
	json.NewEncoder(w).Encode(retrievedZuivel)
}

//WITH QUERY
func handleGetBroodWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat1, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetBroodbelegWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat2, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetFruitWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat3, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetKruidWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat4, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetSnoepWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat5, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetVleesWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat6, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetZuivelWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if !ok || len(sQuery[0]) < 1 || sQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		retrievedProducts = getProductSearch(cat7, sQuery[0])
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}

func handlePostBroodWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	acQuery, ok := r.URL.Query()["ac"]
	pQuery, ok := r.URL.Query()["p"]
	if !ok || len(acQuery[0]) < 1 || acQuery[0] == "0" {
		w.WriteHeader(400)
	} else if !ok || len(pQuery[0]) < 1 || pQuery[0] == "0" {
		w.WriteHeader(400)
	} else {
		postProductAmount(cat1, pQuery[0], acQuery[0])
	}
}
func handlePostBroodbelegWQuery(w http.ResponseWriter, r *http.Request) {

}
