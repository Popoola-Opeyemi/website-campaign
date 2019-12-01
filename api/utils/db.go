package utils

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

// InitDb sets up the database layer
func InitDb(cfg *ini.File, log *zap.Logger) (db *gorm.DB, err error) {
	log.Info("Initializing database connection")

	dbCfg, err := cfg.GetSection("DATABASE")

	if err != nil {
		log.Error(err.Error())
		return
	}

	driver := dbCfg.Key("driver").String()

	DSN := dbCfg.Key("dsn").String()

	db, err = gorm.Open(driver, DSN)

	// make table singular
	db.SingularTable(true)
	db.LogMode(true)

	// Set connection limits for max and open conns
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info(fmt.Sprintf("connected to db with driver %s", driver))

	return
}
