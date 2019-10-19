package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func New(c *Config) (*Database, error) {
	p := adapt(c)

	db, err := gorm.Open(p.DBType, p.DSN)
	if err != nil {
		return nil, err
	}

	if p.Debug {
		db = db.Debug()
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	// 全局禁用表名复数
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(p.MaxIdleConns)
	db.DB().SetMaxOpenConns(p.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(p.MaxLifetime) * time.Second)
	return wrap(db), nil
}

func adapt(c *Config) *param {
	var dsn string
	switch c.Gorm.DBType {
	case "mysql":
		dsn = c.MySQL.dsn()
	case "sqlite3":
		dsn = c.Sqlite3.dsn()
	case "postgres":
		dsn = c.Postgres.dsn()
	}

	return &param{
		Debug:        c.Gorm.Debug,
		DBType:       c.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: c.Gorm.MaxIdleConns,
		MaxLifetime:  c.Gorm.MaxLifetime,
		MaxOpenConns: c.Gorm.MaxOpenConns,
	}
}

func wrap(db *gorm.DB) *Database {
	return &Database{db}
}

type Database struct {
	*gorm.DB
}
