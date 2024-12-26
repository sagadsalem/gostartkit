package db

import (
	"github.com/sajadsalem/startkit/config"
	"github.com/sajadsalem/startkit/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectAndMigrate(config *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&domain.User{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
