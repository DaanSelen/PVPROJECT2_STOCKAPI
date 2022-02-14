package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID          int    `json:"id"`
	Naam        string `json:"naam"`
	Hoeveelheid int    `json:"hoeveelheid"`
	Status      string `json:"status"`
}

var ( //ALL GLOBAL VARIABLES
	requestCounter int = 0

	databaseUser     string
	databaseName     string
	databasePort     string
	databaseIp       string
	databasePassword string

	voorraad *sql.DB
	err      error
)

const ( //ALL GLOBAL CONSTANTS
	errorTag   string = "[Error]"
	infoTag    string = "[Info]"
	warningTag string = "[Warning]"

	cat1 string = "brood"
	cat2 string = "broodbeleg"
	cat3 string = "fruit"
	cat4 string = "kruid"
	cat5 string = "snoep"
	cat6 string = "vlees"
	cat7 string = "zuivel"
)

func getInfoFromConfig(keyword string) string {
	f, err := os.Open("config.txt")
	handleError(err, "Opening config.txt file, perhaps there is no config.txt?")
	defer f.Close()
	lineByLine := bufio.NewScanner(f)
	for lineByLine.Scan() {
		if !strings.Contains(lineByLine.Text(), "#") || lineByLine.Text() != "" { //Skipping empty rows and commented rows
			if strings.Contains(lineByLine.Text(), (keyword + " = ")) {
				info := strings.ReplaceAll(lineByLine.Text(), (keyword + " = "), "")
				return info
			}
		}
	}
}

func initSys() {
	initVars()
	initDBConn()
}

func initVars() {
	databaseUser = getInfoFromConfig("DatabaseUser")
	databaseName = getInfoFromConfig("DatabaseName")
	databasePort = getInfoFromConfig("DatabasePort")
	databaseIp = getInfoFromConfig("DatabaseIp")
	databasePassword = getInfoFromConfig("DatabasePassword")
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

func closeApp() {
	os.Exit(0)
}
