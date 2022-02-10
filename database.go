package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDBConn(password string) {
	voorraad, err = sql.Open("mysql", databaseUser+":"+password+"@tcp("+ipAddress+":"+port+")/"+databaseName) //Connect to database using specified credentials
	if err != nil {
		handleError(err, "Initialising database connection")
	} else {
		testConnection()
	}
}

func testConnection() {
	log.Println(infoTag, "TESTING DATABASE CONNECTION")
	_, err := voorraad.Query("SELECT * FROM brood")
	if err != nil {
		log.Println(errorTag, "DATABASE CONNECTION FAILED. If you don't know what to do, contact an admin!")
		fmt.Scanln()
		closeApp()
	} else {
		log.Println(infoTag, "DATABASE CONNECTION SUCCES.")
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
