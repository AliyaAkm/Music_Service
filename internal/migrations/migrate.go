package migrations

import (
	"PracticeCrud/internal/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func Run(cfg config.DbConfig) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal("Error creating the migrator:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration application error:", err)
	}

	log.Println("Migrations have been successfully applied")
}
