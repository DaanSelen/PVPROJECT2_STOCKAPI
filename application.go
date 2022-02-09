package main

import (
	"database/sql"
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

var (
	requestCounter int = 0

	voorraad *sql.DB
	err      error
)

const ( //Tags
	errorTag   string = "[Error]"
	infoTag    string = "[Info]"
	warningTag string = "[Warning]"

	databaseUser string = "api"
	databaseName string = "voorraad"
	ipAddress    string = "192.168.178.20"
	port         string = "3306"

	cat1 string = "brood"
	cat2 string = "broodbeleg"
	cat3 string = "fruit"
	cat4 string = "kruid"
	cat5 string = "snoep"
	cat6 string = "vlees"
	cat7 string = "zuivel"
)

func initSys() {
	dbpass := getPasswords()
	initDBConn(dbpass)
}

func handleError(err error, location string) {
	if err != nil {
		log.Println(errorTag, "Error encountered at:", location, "error:", err)
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

func getAllTables() []string {
	requestCounter++
	tableList := retrieveAllTables()
	return tableList
}

func getAllProducts(table string) []Product {
	requestCounter++
	retrievedProducts := retrieveAllProducts(("SELECT * FROM " + table))
	return retrievedProducts
}

func getProductSearch(table, query string) []Product {
	requestCounter++
	retrievedProducts := retrieveAllProducts(("SELECT * FROM " + table + " WHERE naam LIKE '%" + query + "%'"))
	return retrievedProducts
}

func postProductAmount(table, product, amountChange string) {
	requestCounter++
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
	requestCounter++
	changeProductAttribute(("UPDATE " + table + " SET Status='" + statusChange + "' WHERE Naam LIKE '%" + product + "%'"))
}
