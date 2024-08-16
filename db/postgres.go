package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/model"
)

type Postgres struct{}


func NewDB() *Postgres {
	return &Postgres{}
}


func (p *Postgres) Connect(creds *model.Credential) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port,
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return dbConn, nil
}


func (p *Postgres) Reset(db *gorm.DB, table string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE " + table + " RESTART IDENTITY").Error; err != nil {
			return err
		}

		return nil
	})
}
