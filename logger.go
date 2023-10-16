package db

import (
	log2 "log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

var defaultLogger = logger.New(
	log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,  // Slow SQL threshold
		LogLevel:                  logger.Error, // Log level
		IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,        // Include params in the SQL log
		Colorful:                  true,         // Enable color
	},
)
