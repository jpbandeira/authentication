package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/jp/authentication/internal/repository/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var errNoModelDefined = errors.New("no model defined")

func defaultGormLogger() gormlogger.Interface {
	return gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormlogger.Config{
			SlowThreshold:             1 * time.Second,
			LogLevel:                  gormlogger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
}

func databaseModels() []any {
	return []any{
		model.User{},
	}
}

func Connect(ctx context.Context) (*gorm.DB, error) {
	_ = godotenv.Load()

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
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	if len(databaseModels()) == 0 {
		return nil, errNoModelDefined
	}

	gormLogger := defaultGormLogger()
	gormConfig := gorm.Config{Logger: gormLogger}

	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(databaseModels()...)
	if err != nil {
		return nil, err
	}

	return db, nil
}
