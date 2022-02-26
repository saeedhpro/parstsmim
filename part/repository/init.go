package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type dbs struct {
	MongoDB *mongo.Client
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

func mongoConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/parts?readPreference=primary&ssl=false"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DBS.MongoDB = client
	DBS.Context = &ctx
	fmt.Println("Successfully connected and pinged.")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err.Error(), "ping error")
	}
}
