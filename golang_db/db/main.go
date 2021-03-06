package main

import (
	"database/sql"
	"fmt"
	"golang_db/db/query"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "12345678"
	NETWORK  = "tcp"
	SERVER   = "172.16.30.10"
	PORT     = 3306
	DATABASE = "arksdb"
)

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("Connection Fail: ", err)
		return
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection Error: ", err.Error())
		return
	}

	fmt.Println("Connection Success!")
	defer db.Close()
	query.Select(db, 12345678)

}
