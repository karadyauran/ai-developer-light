package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"./config"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := config.Config.DBUser + ":" + config.Config.DBPassword + "@tcp(" + config.Config.DBHost + ")/" + config.Config.DBName
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established")
}