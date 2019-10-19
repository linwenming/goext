package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"
)

var (
	Cfg  *Config
	once sync.Once
)

// Config 配置参数
type param struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// Config 配置参数
type Config struct {
	Gorm     Gorm     `toml:"gorm"`
	MySQL    MySQL    `toml:"mysql"`
	Postgres Postgres `toml:"postgres"`
	Sqlite3  Sqlite3  `toml:"sqlite3"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool   `toml:"debug"`
	DBType       string `toml:"db_type"`
	MaxLifetime  int    `toml:"max_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	TablePrefix  string `toml:"table_prefix"`
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DbName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

// DSN 数据库连接串
func (a MySQL) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DbName, a.Parameters)
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
}

// DSN 数据库连接串
func (a Postgres) dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password)
}

// Sqlite3 sqlite3配置参数
type Sqlite3 struct {
	Path string `toml:"path"`
}

// DSN 数据库连接串
func (a Sqlite3) dsn() string {
	return a.Path
}

func LoadDbConfig(fpath string) (*Config, error) {
	c, err := parseConfig(fpath)
	if err != nil {
		return nil, err
	}
	Cfg = c
	return c, nil
}

// ParseConfig 解析配置文件
func parseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
