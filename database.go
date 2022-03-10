package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	voorraad *sql.DB
	err      error
)

func initAndTestDBConn() {
	voorraad, err = sql.Open("mysql", configData[1]+":"+configData[2]+"@tcp("+configData[3]+":"+configData[4]+")/"+configData[5]) //Connect to database using specified credentials
	if err != nil {
		handleError(err, "Initialising database connection")
	} else {
		testConnection()
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

func retrieveMaxID(tableName string) int {
	data, err := voorraad.Query("SELECT MAX(ID) FROM " + tableName)
	handleError(err, "Sending SQL Query for Max ID")
	defer data.Close()
	var maxID int
	for data.Next() {
		data.Scan(&maxID)
	}
	return maxID
}
