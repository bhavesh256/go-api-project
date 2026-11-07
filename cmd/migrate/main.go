package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bhavesh256/go-api-project/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustLoad()

	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down>")
	}

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl,
	)

	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("no change")
				return
			}
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("no change")
				return
			}
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("Migration completed successfully")
}