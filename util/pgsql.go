package util

import (
	"Server/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB(dbConfig conf.DatabaseConfig) (*gorm.DB, error) {
	dsn := dbConfig.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}
