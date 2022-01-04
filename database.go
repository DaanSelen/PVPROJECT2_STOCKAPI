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

func retrieveKruidenAmount() []Product {
	data, err := voorraad.Query("SELECT * FROM kruiden")
	handleError(err, "Sending SQL Query for all kruiden")
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

func retrieveAllProducts(product string) []Product {
	data, err := voorraad.Query("SELECT * FROM " + product)
	handleError(err, ("Sending SQL Query for all " + product))
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
