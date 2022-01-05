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

func getAllProducts(table string) []Product {
	retrievedProducts := retrieveAllProducts(("SELECT * FROM " + table))
	return retrievedProducts
}

func getProductSearch(table, query string) []Product {
	retrievedProducts := retrieveAllProducts(("SELECT * FROM " + table + " WHERE naam LIKE '%" + query + "%'"))
	return retrievedProducts
}

func postProductAmount(table, product, amountChange string) {
	giveProductAmount(("UPDATE " + table + " SET Hoeveelheid=Hoeveelheid+" + amountChange + " WHERE Naam='" + product + "'"))
}
