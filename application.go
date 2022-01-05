package main

import (
	"io/ioutil"
	"log"
)

type Product struct {
	ID          int    `json:"id"`
	Naam        string `json:"naam"`
	Hoeveelheid int    `json:"hoeveelheid"`
	Status      string `json:"status"`
}

func initSys() {
	dbpass := getPasswords()
	initDBConn(dbpass)
}

func handleError(err error, location string) {
	if err != nil {
		log.Println("Error encountered at:", location, "error:", err)
	}
}

func getPasswords() string {
	data, err := ioutil.ReadFile("dbpass.key")
	handleError(err, "Retrieving Passwords")
	dbpassword := string(data)
	return dbpassword
}

func getAllProducts(product string) []Product {
	retrievedProducts := retrieveAllProducts(("SELECT * FROM " + product))
	return retrievedProducts
}

func getProductSearch(product, query string) []Product {
	retrievedProducts := retrieveAllProducts(("SELECT * FROM " + product + " WHERE naam LIKE '%" + query + "%'"))
	return retrievedProducts
}
