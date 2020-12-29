package database

import "gorm.io/gorm"

var DB *gorm.DB

type Config struct {
	ConnectionString string
}