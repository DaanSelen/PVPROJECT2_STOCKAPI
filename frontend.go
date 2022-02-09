package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println(infoTag, "CELDSERV API APPLICATION INITIALISING")
	initSys()

	myRouter := mux.NewRouter().StrictSlash(true)
	log.Println(infoTag, "READY FOR REQUESTS")

	myRouter.HandleFunc("/", handleVoorraadRoot).Methods("GET")
	myRouter.HandleFunc("/voorraad", handleVoorraadRoot).Methods("GET")
	myRouter.HandleFunc("/requests", handleRequestCounter).Methods("GET")

	myRouter.HandleFunc("/voorraad/all/brood", handleGetAllBrood).Methods("GET")
	myRouter.HandleFunc("/voorraad/brood", handleGetBroodWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/brood", handlePostBroodWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/all/broodbeleg", handleGetAllBroodbeleg).Methods("GET")
	myRouter.HandleFunc("/voorraad/broodbeleg", handleGetBroodbelegWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/broodbeleg", handlePostBroodbelegWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/all/fruit", handleGetAllFruit).Methods("GET")
	myRouter.HandleFunc("/voorraad/fruit", handleGetFruitWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/fruit", handlePostFruitWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/all/kruid", handleGetAllKruid).Methods("GET")
	myRouter.HandleFunc("/voorraad/kruid", handleGetKruidWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/kruid", handlePostKruidWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/all/snoep", handleGetAllSnoep).Methods("GET")
	myRouter.HandleFunc("/voorraad/snoep", handleGetSnoepWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/snoep", handlePostSnoepWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/all/vlees", handleGetAllVlees).Methods("GET")
	myRouter.HandleFunc("/voorraad/vlees", handleGetVleesWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/vlees", handlePostVleesWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/all/zuivel", handleGetAllZuivel).Methods("GET")
	myRouter.HandleFunc("/voorraad/zuivel", handleGetZuivelWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/zuivel", handlePostZuivelWQuery).Methods("PATCH")

	http.ListenAndServe(":61909", myRouter)
}

//ROOT RESPONSE
func handleVoorraadRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Root directory for the /voorraad tree. Available options:")
	tableList := getAllTables()
	json.NewEncoder(w).Encode(tableList)
}

func handleRequestCounter(w http.ResponseWriter, r *http.Request) {
	log.Println(infoTag, "Current amount of request(s):", requestCounter)
	json.NewEncoder(w).Encode("Current amount of request(s) (excluding this one): " + changeToString(requestCounter))
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
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat1, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetBroodbelegWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat2, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetFruitWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat3, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetKruidWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat4, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetSnoepWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat5, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetVleesWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat6, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}
func handleGetZuivelWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery[0]) > 1 || sQuery[0] != "0" {
		retrievedProducts = getProductSearch(cat7, sQuery[0])
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(retrievedProducts)
}

func handlePostBroodWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat1, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat1, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePostBroodbelegWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat2, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat2, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePostFruitWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat3, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat3, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePostKruidWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat4, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat4, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePostSnoepWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat5, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat5, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePostVleesWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat6, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat6, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePostZuivelWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat7, pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat7, pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
