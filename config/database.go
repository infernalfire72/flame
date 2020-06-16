package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB

type DatabaseConfig struct {
	Host		string
	Database	string
	Username	string
	Password	string
}

func (c *DatabaseConfig) String() string {
	return fmt.Sprintf("%s:%s@%s/%s", c.Username, c.Password, c.Host, c.Database)
}