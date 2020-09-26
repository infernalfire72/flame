package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/infernalfire72/flame/bancho"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/osuapi"
	"github.com/infernalfire72/flame/web"
	"github.com/jmoiron/sqlx"
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

	osuapi.Enabled = conf.OsuApi.Enable
	osuapi.Key = conf.OsuApi.Key

	web.Start(&conf.Web)
	bancho.Start(&conf.Bancho)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-exit

	config.Database.Close()

	log.Info("Exiting...")
}
