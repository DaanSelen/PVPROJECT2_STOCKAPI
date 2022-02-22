package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	voorraad *sql.DB
	err      error
)

func initDBConn() {
	voorraad, err = sql.Open("mysql", configData[1]+":"+configData[2]+"@tcp("+configData[3]+":"+configData[4]+")/"+configData[5]) //Connect to database using specified credentials
	if err != nil {
		handleError(err, "Initialising database connection")
	} else {
		testConnection()
	}
}

func testConnection() {
	log.Println(infoTag, "TESTING DATABASE CONNECTION")
	_, err := voorraad.Query("SHOW Databases")
	if err != nil {
		log.Println(errorTag, "DATABASE CONNECTION FAILED. If you don't know what to do, contact an admin!")
		closeApp()
	} else {
		log.Println(infoTag, "DATABASE SERVER AND DATABASE CONNECTION SUCCES, TESTING DATABASE TABLES.")
		for x := range cat {
			_, err := voorraad.Query("SELECT * from " + cat[x])
			if err != nil {
				log.Println(infoTag, "NO TABLE CALLED '"+cat[x]+"' DETECTED, CREATING IT.")
				_, err := voorraad.Query("CREATE TABLE " + cat[x] + " (ID INT(20) NOT NULL, Naam VARCHAR(50) NOT NULL, Hoeveelheid INT(50) NOT NULL, Status VARCHAR(100) NOT NULL)")
				handleError(err, "Creating "+cat[x]+" table")
			}
		}
	}
}

func retrieveAllTables() []string {
	var tableList []string

	data, err := voorraad.Query("SHOW tables")
	handleError(err, "Getting tables from voorraad")
	defer data.Close()
	for data.Next() {
		var table string
		data.Scan(&table)
		tableList = append(tableList, table)
	}
	return tableList
}

func retrieveHoeveelheid(table, product string) int {
	var hoeveelheid int
	queryString := "SELECT Hoeveelheid FROM " + table + " WHERE Naam LIKE '%" + product + "%'"

	data, err := voorraad.Query(queryString)
	handleError(err, "Getting hoeveelheid Query")
	defer data.Close()
	for data.Next() {
		data.Scan(&hoeveelheid)
	}
	return hoeveelheid
}

func retrieveAllProducts(queryString string) []Product {
	data, err := voorraad.Query(queryString)
	handleError(err, "Sending SQL Query for all ")
	defer data.Close()
	var products []Product
	for data.Next() {
		var product Product
		err := data.Scan(&product.ID, &product.Naam, &product.Hoeveelheid, &product.Status)
		handleError(err, "Scanning the content of the Database response")
		products = append(products, product)
	}
	return products
}

func changeProductAttribute(queryString string) {
	_, err := voorraad.Query(queryString)
	handleError(err, "Sending SQL Query for changing product amount(hoeveelheid)")
}
