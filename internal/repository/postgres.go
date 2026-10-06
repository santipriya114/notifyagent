package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


func ConnectDB() {

	connStr := `
	host=notificationdb12345.postgres.database.azure.com
	port=5432
	user=postgresadmin
	password=Moni@1429
	dbname=notificationdb
	sslmode=require
	connect_timeout=60
	`

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Connected to Azure PostgreSQL")
}