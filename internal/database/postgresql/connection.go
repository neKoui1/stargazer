package postgresql

import (
	"fmt"
	"log"
	"stargazer/internal/config"
	"stargazer/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Postgre *gorm.DB

func ConnectPostgresql(cfg *config.DatabaseConfig) error {
	dsn := cfg.GetDSN()

	var logLevel logger.LogLevel
	if cfg.SSLMode == "disable" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	db, err := gorm.Open(
		postgres.Open(dsn),
		gormConfig,
	)
	if err != nil {
		return fmt.Errorf(
			"database: failed to connect to postgresql: %w",
			err,
		)
	}

	// 获取底层的sql.DB对象
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf(
			"database: failed to get sql.DB: %w",
			err,
		)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 1)

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf(
			"database: failed to ping postgresql database: %w",
			err,
		)
	}

	Postgre = db
	log.Println(
		"database: connected to postgresql database successfully",
	)
	return nil
}

func AutoMigrate() error {
	if Postgre == nil {
		return fmt.Errorf(
			"database: failed to connect to postgresql database",
		)
	}

	// 自动迁移所有模型
	err := Postgre.AutoMigrate(
		&models.User{},
		// other models
	)

	if err != nil {
		return fmt.Errorf(
			"database: failed to auto migrate: %w",
			err,
		)
	}

	log.Println("database: auto migration completed successfully")
	return nil
}

func Close() error {
	if Postgre == nil {
		return nil
	}

	sqlDB, err := Postgre.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func GetPostgre() *gorm.DB {
	return Postgre
}

func HealthCheck() error {
	if Postgre == nil {
		return fmt.Errorf(
			"database: failed to connect to postgresql database",
		)
	}

	sqlDB, err := Postgre.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
