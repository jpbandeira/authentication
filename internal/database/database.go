package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect(ctx context.Context) (*sql.DB, error) {
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if user == "" {
		user = "postgres"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if dbName == "" {
		dbName = "auth_db"
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Erro ao abrir conexão: %v", err)
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		log.Printf("Erro ao pingar o banco: %v", err)
		return nil, err
	}

	return db, nil
}
