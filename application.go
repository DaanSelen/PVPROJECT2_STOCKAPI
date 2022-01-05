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
	hoeveelheid := retrieveHoeveelheid(("SELECT Hoeveelheid FROM " + table))
	intAmount := hoeveelheid + changeToInt(amountChange)
	giveProductAmount(("UPDATE " + table + " SET Hoeveelheid=" + changeToString(intAmount) + " WHERE Naam='" + product + "'"))
}
