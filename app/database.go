package app

import (
	"github.com/andriimwks/go-fiber-template/config"
	"github.com/andriimwks/go-fiber-template/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.SQLite.Path), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
