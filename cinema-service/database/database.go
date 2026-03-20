package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"cinema/config"
	"cinema/internal/app/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB      *pgxpool.Pool
	Queries *repository.Queries
	once    sync.Once
)

func InitDB() {
	once.Do(func() {
		dbConfig, err := config.GetDatabaseConfig()
		if err != nil {
			log.Fatalf("Failed to get database config: %v", err)
		}

		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

		poolConfig, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Unable to parse database config DSN: %v", err)
		}

		poolConfig.MaxConns = 20
		poolConfig.MinConns = 5
		poolConfig.MaxConnIdleTime = 5 * time.Minute
		poolConfig.MaxConnLifetime = time.Hour
		poolConfig.HealthCheckPeriod = time.Minute

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		DB, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

		if err = DB.Ping(ctx); err != nil {
			log.Fatalf("Could not ping database: %v", err)
		}

		Queries = repository.New(DB)
		log.Println("Database connection successfully initialized!")
	})
}

func Fatal(err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		log.Printf("Non-PG Error: %v", err)
		return errors.New("internal server error")
	}

	// 1. Detailed log for developer
	log.Printf("[DB ERROR] Code: %s | Msg: %s | Detail: %s | Constraint: %s",
		pgErr.Code, pgErr.Message, pgErr.Detail, pgErr.ConstraintName)

	// 2. Concise English message for user/client
	var msg string
	category := pgErr.Code[:2]

	switch category {
	case "23": // Constraints
		switch pgErr.Code {
		case "23505":
			msg = "data already exists"
		case "23503":
			msg = "related data not found or still in use"
		case "23502":
			msg = "required field is missing"
		default:
			msg = "data integrity violation"
		}

	case "08": // Connection
		msg = "database connection failed"

	case "28", "42": // Auth / Syntax
		msg = "database configuration error"

	case "57": // Shutdown/Cancel
		msg = "request canceled by database"

	default:
		msg = fmt.Sprintf("database error (%s)", pgErr.Code)
	}

	return errors.New(msg)
}

func CloseDB() {
	if DB != nil {
		log.Println("Closing database connection pool.")
		DB.Close()
	}
}
