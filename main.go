package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/infernalfire72/flame/bancho"
	"github.com/infernalfire72/flame/cache"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/web"
	"github.com/jmoiron/sqlx"
)

func init() {
	cache.Init()
}

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

	go bancho.Start(&conf.Bancho)
	web.Start(&conf.Web)
}
