package main

import (
	"bufio"
	"fmt"
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

	cat                     = [7]string{"brood", "broodbeleg", "fruit", "kruid", "snacks", "vlees", "zuivel"}
	configKeywords          = []string{"ApiPort", "DatabaseUser", "DatabasePassword", "DatabaseIP", "DatabasePort", "DatabaseName"}
	configData     []string //INDEXES MUST MATCH configKeywords corresponding INDEX, OTHERWISE THE PROGRAM WILL FAIL.
	productData    [][]string
)

const ( //ALL GLOBAL CONSTANTS
	errorTag   string = "[Error]"
	infoTag    string = "[Info]"
	warningTag string = "[Warning]"
)

func initSys() {
	initVars()
	initProducts()
	initAndTestDBConn()
}

func initVars() {
	for x := range configKeywords {
		y := getInfoFromConfig(configKeywords[x])
		configData = append(configData, y)
	}
}

func initProducts() {
	for x := range cat {
		getInfoFromProducts(cat[x])
	}
}

func getInfoFromConfig(keyword string) string {
	var info string
	f, err := os.Open("config.txt")
	handleError(err, "Opening config.txt file, perhaps there is no config.txt?")
	defer f.Close()
	lineByLine := bufio.NewScanner(f)
	for lineByLine.Scan() {
		if !strings.Contains(lineByLine.Text(), "#") || lineByLine.Text() != "" { //Skipping empty rows and commented rows
			if strings.Contains(lineByLine.Text(), (keyword + " = ")) {
				info = strings.ReplaceAll(lineByLine.Text(), (keyword + " = "), "")
			}
		}
	}
	return info
}

func getInfoFromProducts(keyword string) {
	f, err := os.Open("products.txt")
	handleError(err, "Opening products.txt file, perhaps there is no config.txt?")
	defer f.Close()
	lineByLine := bufio.NewScanner(f)
	for lineByLine.Scan() {
		if !strings.Contains(lineByLine.Text(), "#") || lineByLine.Text() != "" { //Skipping empty rows and commented rows
			if lineByLine.Text() == ("CATEGORY " + keyword) {
				lineByLine.Scan()
				localProducts := strings.Split(lineByLine.Text(), ", ")
				productData = append(productData, localProducts)
			}
		}
	}
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

func askDBQuestion() bool {
	log.Println(infoTag, "IS THE GIVEN DATABASE SETUP CORRECTLY? (SELECT no IF THE DATABASE IS EMPTY OR INCORRECTLY SETUP) [Y/n]")
	var userInput string
	fmt.Scanln(&userInput)
	return checkDBInput(strings.ToLower(userInput))
}
func checkDBInput(userInput string) bool {
	if userInput == "" || userInput == "y" {
		return true
	} else {
		return false
	}
}

func closeApp() {
	fmt.Scanln()
	os.Exit(0)
}
