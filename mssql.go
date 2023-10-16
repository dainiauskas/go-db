package db

import (
	"fmt"

	"github.com/dainiauskas/go-config"
	"github.com/dainiauskas/go-log"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Ms *gorm.DB

func NewMsSQL(cfg *config.Database) error {
	log.Info("Connecting to MsSQL..")

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;collation=%s;",
		cfg.Host, cfg.User, cfg.Pass, cfg.Name, cfg.Collation)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: defaultLogger,
	})
	if err != nil {
		return err
	}

	Ms = db

	log.Trace("Connected to DB %s...", cfg.Name)

	return nil
}
