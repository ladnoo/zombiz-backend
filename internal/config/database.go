package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=postgres password=9996 dbname=zombiz sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("не удалось подключиться к базе данных:", err)
	}

	log.Println("Успешное подключение к базе данных PostgreSQL")

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
}
