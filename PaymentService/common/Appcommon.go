package common

import (
	. "PaymentService/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbcon *sql.DB
var config = Config{}

func init() {
	fmt.Println("Init service ....")
	config.Read()
	initLogger(config)
	if dbcon == nil {
		dbcon = initConnectionPool(config)
	}

}

func GetConnection() (db *sql.DB) {
	return dbcon
}

func GetServicePort() string {
	return config.ServicePort
}

func GetUrlService() string {
	return config.UrlService
}
