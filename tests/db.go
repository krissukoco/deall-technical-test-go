package tests

import (
	"fmt"

	"github.com/krissukoco/deall-technical-test-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresTestDsn(dbname string) (string, error) {
	cfg, err := config.Load("test")
	if err != nil {
		return "", err
	}
	dbc := cfg.Database
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbc.Host, dbc.Port, dbc.Username, dbc.Password, dbc.DbName), nil
}

func NewTestDb(dbname string) (*gorm.DB, error) {
	dsn, err := postgresTestDsn(dbname)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
