package pkg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Hostname string
	Port     string
	DbName   string
	Username string
	Password string
	SslMode  string
}

func NewConfig(host, port, dbname, username, password, sslmode string) Config {
	return Config{
		Hostname: host,
		Port:     port,
		DbName:   dbname,
		Username: username,
		Password: password,
		SslMode:  sslmode,
	}
}

func Init(cfg Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%s user =%s dbname=%s password=%s sslmode=%s",
		cfg.Hostname, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SslMode))
	if err != nil {
		return nil, err
	}
	return conn, err
}
