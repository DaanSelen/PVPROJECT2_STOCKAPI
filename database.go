package main

import (
	"database/sql"
	"log"
)

var (
	voorraad *sql.DB
	err      error
)

func initDBConn(password string) {
	_, err := sql.Open("mysql", "daan:"+password+"@tcp(192.168.178.23:3307)/voorraad")
	if err != nil {
		log.Println("Connection to Database Server failed: ", err)
	}
}

func retrieveKruidenAmount() {
	data, err := voorraad.Query("")
	handleError(err)
	defer data.Close()
	for data.Next() {
		err := data.Scan(&ID)
		handleError(err)
	}
}
