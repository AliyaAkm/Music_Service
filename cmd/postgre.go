package main

import (
	"PracticeCrud/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func NewDBConnection(cfg config.DbConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("База данных недоступна:", err)
	}

	return db
}
