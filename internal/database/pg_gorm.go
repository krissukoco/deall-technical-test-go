package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DefaultPostgresPort     = 5432
	DefaultPostgresTimezone = "Asia/Jakarta"
)

func envErr(key string) error {
	return fmt.Errorf("env %s is not set", key)
}

func postgresDsn(host, user, password, dbname, timezone string, port int, enableSsl bool) string {
	sslMode := "disable"
	if enableSsl {
		sslMode = "enable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=%d TimeZone=%s", host, user, password, dbname, sslMode, port, timezone)
}

func NewPostgresGorm(host, user, password, dbname, timezone string, port int, enableSsl bool) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(postgresDsn(host, user, password, dbname, timezone, port, enableSsl)),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDefaultPostgresGorm() (*gorm.DB, error) {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		return nil, envErr("POSTGRES_HOST")
	}
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		return nil, envErr("POSTGRES_USER")
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		return nil, envErr("POSTGRES_PASSWORD")
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		return nil, envErr("POSTGRES_DB")
	}
	port := DefaultPostgresPort
	portStr, exists := os.LookupEnv("POSTGRES_PORT")
	if exists {
		portConv, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		port = portConv
	}
	tz := DefaultPostgresTimezone
	timezone, exists := os.LookupEnv("POSTGRES_TIMEZONE")
	if exists {
		tz = timezone
	}

	enableSsl := false
	sslEnv := os.Getenv("POSTGRES_SSL")
	switch sslEnv {
	case "true", "yes", "enable":
		enableSsl = true
	}
	return NewPostgresGorm(host, user, password, dbname, tz, port, enableSsl)
}
