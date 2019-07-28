package common

import (
	"database/sql"
	"fmt"

	log "github.com/Sirupsen/logrus"

	. "employee/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

func GetdbConnforgorm() (db *gorm.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "test"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	return db
}
