package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "os"
	"strconv"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		dbName = "go_psql_demo_webapp_db"
	}
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		dbUsername = "postgres"
	}
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		dbPassword = ""
	}
	var dbConnectionPoolSize int
	var err error
	dbConnectionPoolSizeStr := os.Getenv("DB_CONNECTION_POOL_SIZE")
	if dbConnectionPoolSizeStr == "" {
		dbConnectionPoolSize = 50
	} else {
		dbConnectionPoolSize, err = strconv.Atoi(dbConnectionPoolSizeStr)
		if err != nil {
			log.Fatalf("Invalid DB_CONNECTION_POOL_SIZE value: %s", dbConnectionPoolSizeStr)
		}
	}
	URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", URL)
	db.SetMaxOpenConns(dbConnectionPoolSize)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database!")
}
