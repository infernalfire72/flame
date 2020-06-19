package main

import (
	"github.com/infernalfire72/flame/bancho"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/web"
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		return
	}

	config.Database, err = sqlx.Open("mysql", conf.Database.String())
	if err != nil {
		log.Error(err)
		return
	}
	if err = config.Database.Ping(); err != nil {
		log.Error(err)
		return
	}

	go web.Start(&conf.Web)
	bancho.Start(&conf.Bancho)
}