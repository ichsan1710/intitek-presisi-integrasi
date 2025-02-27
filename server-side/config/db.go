package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:root1234@/go_products")
	if err != nil {
		panic(err)
	}

	log.Println("Database has been connected")

	DB = db
}
