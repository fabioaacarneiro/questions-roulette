package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("erro ao carregar o dotenv.", err)
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PARAMS"),
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
