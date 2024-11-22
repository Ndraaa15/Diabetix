package config

import (
	"fmt"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(env *env.Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		env.DatabaseHost,
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseName,
		env.DatabasePort,
		env.DatabaseSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	db.Debug()

	if err != nil {
		zap.S().Fatalf("Unable to connect to database: %v", err)
	}

	return db
}
