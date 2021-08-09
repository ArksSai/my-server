package golang_db

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/golang_db/query"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "12345678"
	NETWORK  = "tcp"
	SERVER   = "172.16.30.10"
	//SERVER   = "13.208.168.50"
	PORT     = 3306
	DATABASE = "arksdb"
)

func Connect() string {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("Connection Fail: ", err)
		return "Connection Fail"
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection Error: ", err.Error())
		return "Connection Error"
	}

	fmt.Println("Connection Success!")
	defer db.Close()
	return query.Select(db, 12345678)

}
