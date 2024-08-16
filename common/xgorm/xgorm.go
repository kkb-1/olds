package xgorm

import (
	"application/common/xzap"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config interface {
	getDSN() string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func (cfg Mysql) getDSN() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName)
}

func (cfg Postgres) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DbName,
		cfg.Port)
}

func Open(cfg Config, logger xzap.Log) (*gorm.DB, error) {
	dsn := cfg.getDSN()
	var open gorm.Dialector
	switch cfg.(type) {
	case Mysql:
		open = mysql.Open(dsn)
	case Postgres:
		open = postgres.Open(dsn)
	}

	db, err := gorm.Open(open, &gorm.Config{Logger: logger})
	return db, err
}

func MustOpen(cfg Config, logger xzap.Log) *gorm.DB {
	db, err := Open(cfg, logger)
	if err != nil {
		panic(err)
	}

	return db
}
