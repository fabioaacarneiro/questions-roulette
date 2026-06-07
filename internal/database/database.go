package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("erro ao carregar o dotenv.", err)
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		log.Println("erro ao conectar ao banco de dados.", err)
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		log.Println("erro ao pingar ao banco de dados.", err)
		return nil, err
	}

	return DB, nil
}

func GetDB() *sql.DB {
	return DB
}
