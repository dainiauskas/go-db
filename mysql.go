package db

import (
	"github.com/dainiauskas/go-config"
	"github.com/dainiauskas/go-log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBC struct {
	Orm *gorm.DB
}

var My *DBC

func New(cfg *config.Database) (*DBC, error) {
	log.Info("Connecting to MySql..")

	dsn := cfg.FormatDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: defaultLogger,
	})
	if err != nil {
		return nil, err
	}

	My = &DBC{
		Orm: db,
	}

	log.Trace("Connected to DB %s...", cfg.Name)

	return My, nil
}

func Migrate(models ...interface{}) error {
	return My.Orm.AutoMigrate(
		models...,
	)
}
