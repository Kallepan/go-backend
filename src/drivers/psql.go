package drivers

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

type PSQLConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

func (pc *PSQLConfig) Init(ctx context.Context) {
	once.Do(
		func() {
			connectionString := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
				pc.Host,
				pc.User,
				pc.Password,
				pc.DBName,
				pc.Port,
			)
			instance, err := sql.Open("postgres", connectionString)
			if err != nil {
				slog.Error(err.Error())
				panic(err)
			}

			// Close the DB connection when the context is done
			go func() {
				<-ctx.Done()
				if err := db.Close(); err != nil {
					slog.Error(err.Error())
					panic(err)
				}
			}()

			// Ping the DB to check if the connection is alive
			if err := instance.Ping(); err != nil {
				slog.Error(err.Error())
				panic(err)
			}

			// Set the global DB instance to the local instance
			db = instance
		},
	)
}

func initMigrations() {
	// Run the migrations
	if db == nil {
		slog.Error("DB connection not initialized")
		panic("DB connection not initialized")
	}

	// Run the migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	migrationPath := os.Getenv("MIGRATION_PATH")
	if migrationPath == "" {
		migrationPath = "file://app/migrations"
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, "postgres", driver)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error(err.Error())
		panic(err)
	} else if err == migrate.ErrNoChange {
		slog.Info("No migration to run")
	}

	slog.Info("Migrations completed!")
}

func ConnectToDB() *sql.DB {
	ctx := context.Background()
	/* Set up the connection String using the dbConfig struct */
	pc := PSQLConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}

	// Initialize the DB connection
	pc.Init(ctx)

	// Verify the DB connection
	if db == nil {
		slog.Error("DB connection not initialized")
		panic("DB connection not initialized")
	}

	if err := db.Ping(); err != nil {
		slog.Error("DB connection not initialized")
		panic("DB connection not initialized")
	}

	// Run the migrations
	initMigrations()

	return db
}
