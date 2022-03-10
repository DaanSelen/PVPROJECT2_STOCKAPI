package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println(infoTag, "CELDSERV STOCKY™ APPLICATION INITIALISING")
	initSys()

	myRouter := mux.NewRouter().StrictSlash(true)
	log.Println(infoTag, "READY FOR REQUESTS ON PORT:", configData[0])

	myRouter.HandleFunc("/", handleVoorraadRoot).Methods("GET")
	myRouter.HandleFunc("/voorraad", handleVoorraadRoot).Methods("GET")
	myRouter.HandleFunc("/requests", handleRequestCounter).Methods("GET")

	myRouter.HandleFunc("/voorraad/brood", handleGetBroodWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/brood", handlePatchBroodWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/broodbeleg", handleGetBroodbelegWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/broodbeleg", handlePatchBroodbelegWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/fruit", handleGetFruitWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/fruit", handlePatchFruitWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/kruid", handleGetKruidWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/kruid", handlePatchKruidWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/snoep", handleGetSnoepWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/snoep", handlePatchSnoepWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/vlees", handleGetVleesWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/vlees", handlePatchVleesWQuery).Methods("PATCH")

	myRouter.HandleFunc("/voorraad/zuivel", handleGetZuivelWQuery).Methods("GET")
	myRouter.HandleFunc("/voorraad/zuivel", handlePatchZuivelWQuery).Methods("PATCH")

	http.ListenAndServe((":" + configData[0]), myRouter)
}

//ROOT RESPONSES AND UTILITIES
func handleVoorraadRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Root directory for the /voorraad tree. Available options:")
	tableList := getAllTables()
	json.NewEncoder(w).Encode(tableList)
}

func handleRequestCounter(w http.ResponseWriter, r *http.Request) {
	log.Println(infoTag, "Current amount of requests:", requestCounter)
	json.NewEncoder(w).Encode("Current amount of requests (excluding this one): " + changeToString(requestCounter))
}

//GET REQUESTS
func handleGetBroodWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[0], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[0])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}
func handleGetBroodbelegWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[1], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[1])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}
func handleGetFruitWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[2], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[2])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}
func handleGetKruidWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[3], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[3])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}
func handleGetSnoepWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[4], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[4])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}
func handleGetVleesWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[5], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[5])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}
func handleGetZuivelWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var retrievedProducts []Product

	sQuery, ok := r.URL.Query()["s"]
	if ok || len(sQuery) > 0 {
		retrievedProducts = getProductSearch(cat[6], sQuery[0])
		json.NewEncoder(w).Encode(retrievedProducts)
	} else {
		retrievedBroden := getAllProducts(cat[6])
		json.NewEncoder(w).Encode(retrievedBroden)
	}
}

//PATCH REQUESTS
func handlePatchBroodWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[0], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[0], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePatchBroodbelegWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[1], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[1], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePatchFruitWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[2], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[2], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePatchKruidWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[3], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[3], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePatchSnoepWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[4], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[4], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePatchVleesWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[5], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[5], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
func handlePatchZuivelWQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pQuery, ok1 := r.URL.Query()["p"]
	acQuery, ok2 := r.URL.Query()["ac"]
	scQuery, ok3 := r.URL.Query()["sc"]
	if ok1 || len(pQuery) > 0 {
		if ok2 || len(acQuery) > 0 {
			postProductAmount(cat[6], pQuery[0], acQuery[0])
		} else if ok3 || len(scQuery) > 0 {
			postProductStatus(cat[6], pQuery[0], scQuery[0])
		} else {
			json.NewEncoder(w).Encode("Only the 'p' query has been detected, possible other queries include: ac, sc")
		}
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("The 'p' query is required for this end of the API. Contact the admin if you don't know what to do.")
	}
}
