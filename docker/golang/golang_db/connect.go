package golang_db

import (
	"database/sql"
	"fmt"
	"strconv"

	golang_db "github.com/gin-gonic/golang_db/query"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

//const (
//	USERNAME = "root"
//	PASSWORD = "12345678"
//	NETWORK  = "tcp"
//	//SERVER   = "172.16.30.10"
//	SERVER   = "13.208.191.215"
//	PORT     = 3306
//	DATABASE = "arksdb"
//)

type db_config struct {
	USERNAME string
	PASSWORD string
	NETWORK  string
	SERVER   string
	PORT     int
	DATABASE string
}

func Connect() (*sql.DB, error) {
	config := viper.New()
	config.AddConfigPath("./etc/database")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	db_connect := db_config{}
	db_connect.USERNAME = config.GetString("mysql.master.username")
	db_connect.PASSWORD = config.GetString("mysql.master.password")
	db_connect.NETWORK = config.GetString("mysql.master.network")
	db_connect.SERVER = config.GetString("mysql.master.host")
	db_connect.PORT = config.GetInt("mysql.master.port")
	db_connect.DATABASE = config.GetString("mysql.master.dbname")

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", db_connect.USERNAME, db_connect.PASSWORD, db_connect.NETWORK, db_connect.SERVER, db_connect.PORT, db_connect.DATABASE)
	db, err := sql.Open("mysql", conn)
	return db, err
}

func CheckPassword(id string, password string) bool {
	db, err := Connect()
	if err != nil {
		fmt.Println("Connection Fail: ", err)
		return false
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection Error: ", err.Error())
		return false
	}

	defer db.Close()

	if err != nil {
		fmt.Println("Connection Fail: ", err)
		return false
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection Error: ", err.Error())
		return false
	}
	intid, err := strconv.Atoi(id)
	user_data, err := golang_db.SelectById(db, intid)

	return user_data.Password == password
}
