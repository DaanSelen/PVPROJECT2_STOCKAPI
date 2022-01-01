package main

import (
	"io/ioutil"
	"log"
)

func initSys() {

	dbpass := getPasswords()
	initDBConn(dbpass)
}

func handleError(err error) {
	log.Println("Error encountered, error:", err)
}

func getPasswords() string {
	data, err := ioutil.ReadFile("dbpass.key")
	handleError(err)
	dbpassword := string(data)
	return dbpassword
}

func getKruidenAmount() {
	retrieveKruidenAmount()
}
