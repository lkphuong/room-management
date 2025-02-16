package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func ConnectionSqlServer() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable;",
		databaseHost, databaseUser, databasePassword, databasePort, databaseName)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	fmt.Println("Connected!")

	return db
}

func DynamicConnectionSqlServer(host string, user string, pwd string, port string, name string) *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable;",
		host, user, pwd, port, name)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	fmt.Println("Connected!")

	return db
}
