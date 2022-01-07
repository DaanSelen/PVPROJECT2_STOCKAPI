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

func initDBConn(password string) {
	voorraad, err = sql.Open("mysql", "daan:"+password+"@tcp(192.168.178.23:3307)/voorraad") //IF connected with OPENVPN SERVER.
	if err != nil {
		log.Println("Connection to Database Server failed: ", err)
	}
}

func retrieveAllTables() []string {
	var tableList []string

	data, err := voorraad.Query("SHOW tables")
	handleError(err, "Getting tables from Voorraad")
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
