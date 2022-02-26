package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type dbs struct {
	MySQL   *sql.DB
	Context *context.Context
}

var DBS dbs

func Init() {
	mysqlConnection()
}

func mysqlConnection() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_ROOT_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_NAME"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", uri)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	DBS.MySQL = db
	fmt.Println("MySQL Connected!")
}
