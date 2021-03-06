package initialize

import (
	"log"
	"os"
	"time"
	"trainmate/models"

	"github.com/emvi/logbuch"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func InitORM() {
	// config = &conf.Config{}
	if !config.UseMongo {
		// Set up GORM
		gormLogger := logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Minute,
				Colorful:      false,
				LogLevel:      logger.Silent,
			},
		)

		// Connect to database
		var err error
		db, err = gorm.Open(config.Db.GetDialector(), &gorm.Config{Logger: gormLogger})
		if err != nil {
			panic(err)
		}
		if config.Db.Dialect == "sqlite3" {
			db.Raw("PRAGMA foreign_keys = ON;")
			db.DisableForeignKeyConstraintWhenMigrating = true
		}

		db.AutoMigrate(&models.Task{})
		db.AutoMigrate(&models.Experiment{})
		db.AutoMigrate(&models.Job{})
		db.AutoMigrate(&models.Dataset{})

		if config.IsDev() {
			db = db.Debug()
		}
		sqlDb, err := db.DB()
		sqlDb.SetMaxIdleConns(int(config.Db.MaxConn))
		sqlDb.SetMaxOpenConns(int(config.Db.MaxConn))
		if err != nil {
			logbuch.Error(err.Error())
			logbuch.Fatal("could not connect to database")
		}
		// defer sqlDb.Close()

		// Migrate database schema
		// migrations.Run(db, config)
		// 迁移 schema

	}

}
