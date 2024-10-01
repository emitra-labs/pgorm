package pgorm

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/emitra-labs/common/validator"
	"github.com/sethvargo/go-envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	URL             string `env:"PGORM_URL" validate:"required"`
	MaxIdleConns    int    `env:"PGORM_MAX_IDLE_CONNS, default=10"`
	MaxOpenConns    int    `env:"PGORM_MAX_OPEN_CONNS, default=100"`
	RevealLogValues bool   `env:"PGORM_REVEAL_LOG_VALUES, default=false"`
}

var DB *gorm.DB

func Open() {
	var config Config

	// Load config from environment variables
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		panic(err)
	}

	// Validate config
	err = validator.Validate(config)
	if err != nil {
		panic(err)
	}

	// Open database connection with customized logging
	DB, err = gorm.Open(postgres.Open(config.URL), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      config.RevealLogValues,
			},
		),
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	// Set database connection pool settings
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Close() error {
	sqlDB, _ := DB.DB()
	return sqlDB.Close()
}
