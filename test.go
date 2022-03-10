package main

import "log"

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
	emptyTables()
	log.Println(infoTag, "INSERTING PRODUCTS")
	for a, b := range productData {
		log.Println(infoTag, "INSERTING INTO TABLE:", cat[a])
		for c := range b {
			_, err := voorraad.Query(("INSERT INTO " + cat[a] + " (ID, Naam, Hoeveelheid, Status) VALUES (" + changeToString(retrieveMaxID(cat[a])+1) + ", '" + productData[a][c] + "', 0, 'N.v.t');"))
			handleError(err, "creating table items")
		}
	}
	log.Println(infoTag, "INSERTING FINISHED")
}

func emptyTables() {
	log.Println(infoTag, "STARTED CLEANING TABLES")
	for x := range cat {
		voorraad.Query("DELETE FROM " + cat[x])
	}
	log.Println(infoTag, "FINISHED CLEANING TABLES")
}

func integrityCheck() bool {
	if x := testTables(); x == len(cat) {
		log.Println(infoTag, "DATABASE INTEGRITY VALIDATED")
		return true
	} else {
		return integrityCheck() //True
	}
}
