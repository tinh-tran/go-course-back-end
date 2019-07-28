package common

import (
	. "PaymentService/config"
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

func initConnectionPool(config Config) (db *sql.DB) {
	log.Info("initConnectionPool ...")
	if dbcon == nil {
		dbDriver := "mysql"
		dbUser := config.DbUser
		dbPass := config.DbPassword
		dbName := config.Database
		dbServer := config.DBServer
		dbPort := config.DBPort
		var err error
		dbcon, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbServer+":"+dbPort+")/"+dbName)
		if err != nil {
			panic(err.Error())
		}
		dbcon.SetMaxOpenConns(100)
		dbcon.SetMaxIdleConns(10)
	}
	return dbcon
}
