package main

import (
	"fmt"
	"github.com/thisisthemurph/ledger-system-shared/database"
	"os"

	"github.com/joho/godotenv"
	"github.com/thisisthemurph/ledger-system-shared/internal/migrator"
)

const MigrationsPath = "file://internal/migrations"

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintln(os.Stdout, "usage: migrate [direction: up|down]")
		os.Exit(1)
	}

	directionArg := os.Args[len(os.Args)-1]
	var direction migrator.MigrationDirection
	switch directionArg {
	case "up":
		direction = migrator.MigrationDirectionUp
	case "down":
		direction = migrator.MigrationDirectionDown
	default:
		_, _ = fmt.Fprintln(os.Stdout, "usage: migrate [direction: up|down]")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if err := run(direction, dbName, dbConnectionString); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func run(direction migrator.MigrationDirection, dbName, connectionString string) error {
	db, err := database.Connect(connectionString)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}
	defer db.Close()

	m := migrator.NewPostgresMigrator(db, dbName, MigrationsPath)
	return m.Migrate(direction)
}
