package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"monegoo/models"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {

	var db_user = os.Getenv("DB_USER")
	var db_host = os.Getenv("DB_HOST")
	var db_name = os.Getenv("DB_NAME")
	var db_password = os.Getenv("DB_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_host, db_name))

	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetCurrencies() []models.Currency {
	var currency []models.Currency
	_, err := dbmap.Select(&currency, "select * from currencies order by id desc limit 1")
	if err == nil {
		return currency

	} else {
		return currency

	}
}
