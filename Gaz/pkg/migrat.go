package pkg

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"time"
)

func Migrat(host, port, dbname, user, password string) error {
	time.Sleep(time.Second)
	m, err := migrate.New(
		"file://migrations", fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		return err
	}
	return m.Up()
}
