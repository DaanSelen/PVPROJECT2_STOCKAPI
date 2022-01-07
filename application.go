package main

import (
	"io/ioutil"
	"log"
	"strconv"
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

func changeToInt(value1 string) int {
	changedValue, err := strconv.Atoi(value1)
	handleError(err, "Converting to int")
	return changedValue
}

func changeToString(value1 int) string {
	changedValue := strconv.Itoa(value1)
	return changedValue
}

func checkIfInt(value1 string) bool {
	if _, err := strconv.Atoi(value1); err == nil {
		return true
	} else {
		return false
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
	amountChangeInt := changeToInt(amountChange)
	hoeveelheidInt := retrieveHoeveelheid(table, product)
	newHoeveelheid := hoeveelheidInt + amountChangeInt

	changeProductAttribute(("UPDATE " + table + " SET Hoeveelheid=" + changeToString(newHoeveelheid) + " WHERE Naam LIKE '%" + product + "%'"))
	changedHoeveelheid := retrieveHoeveelheid(table, product)
	if changedHoeveelheid < 0 {
		changeProductAttribute(("UPDATE " + table + " SET Hoeveelheid=0 WHERE Naam LIKE '%" + product + "%'"))
	}
}

func postProductStatus(table, product, statusChange string) {
	changeProductAttribute(("UPDATE " + table + " SET Status='" + statusChange + "' WHERE Naam LIKE '%" + product + "%'"))
}
