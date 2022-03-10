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

func initAndTestDBConn() {
	voorraad, err = sql.Open("mysql", configData[1]+":"+configData[2]+"@tcp("+configData[3]+":"+configData[4]+")/"+configData[5]) //Connect to database using specified credentials
	if err != nil {
		handleError(err, "Initialising database connection")
	} else {
		testConnection()
	}
}

func testConnection() {
	answer := askDBQuestion()
	log.Println(infoTag, "TESTING DATABASE CONNECTION")
	_, err := voorraad.Query("SHOW Databases")
	if err != nil {
		log.Println(errorTag, "DATABASE CONNECTION FAILED. If you don't know what to do, contact an admin!")
		closeApp()
	} else {
		log.Println(infoTag, "DATABASE SERVER AND DATABASE CONNECTION SUCCES, TESTING DATABASE TABLES.")
		if answer == false {
			if integrityCheck() {
				createTableItems()
			}
		} else {
			_ = integrityCheck()
		}
	}
}

func integrityCheck() bool {
	if x := testTables(); x == len(cat) {
		log.Println(infoTag, "DATABASE INTEGRITY VALIDATED")
		return true
	} else {
		return integrityCheck() //true
	}
}

func testTables() int {
	var grade int = 0
	for x := range cat {
		_, err := voorraad.Query("SELECT * from " + cat[x])
		if err != nil {
			createMissingTable(cat[x])
		} else {
			grade++
		}
	}
	return grade
}

func createMissingTable(keyword string) {
	log.Println(infoTag, "CREATING TABLE:", keyword)
	_, err := voorraad.Query("CREATE TABLE " + keyword + " (ID INT(20) NOT NULL, Naam VARCHAR(50) NOT NULL, Hoeveelheid INT(50) NOT NULL, Status VARCHAR(100) NOT NULL)")
	handleError(err, "Creating "+keyword+" table")
}

func createTableItems() {
	log.Println(infoTag, "INSERTING PRODUCTS")
	for a, b := range productData {
		for c := range b {
			_, err := voorraad.Query(("INSERT INTO " + cat[a] + " (ID, Naam, Hoeveelheid, Status) VALUES (" + changeToString(retrieveMaxID(cat[a])+1) + ", '" + productData[a][c] + "', 0, 'N.v.t');"))
			handleError(err, "creating table items")
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
