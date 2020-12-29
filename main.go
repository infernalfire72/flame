package main

import (
	"github.com/infernalfire72/flame/cache/clans"
	"github.com/infernalfire72/flame/cache/users/stats"
	"github.com/infernalfire72/flame/layouts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"

	"github.com/infernalfire72/flame/bancho"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/osuapi"
	"github.com/infernalfire72/flame/web"
)

func init() {
	conf, err := config.Load()
	if err != nil {
		log.Error(err)
		return
	}

	db, err := gorm.Open(mysql.Open(conf.Database.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Error(err)
		os.Exit(1)
		return
	}

	database.DB = db

	db.AutoMigrate(&layouts.User{})
	db.AutoMigrate(&layouts.UserRelationship{})
	db.AutoMigrate(&layouts.UserSettings{})
	db.AutoMigrate(&clans.Clan{})
	db.AutoMigrate(&layouts.Beatmap{})
	db.AutoMigrate(&layouts.Score{})
	db.Table("scores_relax").AutoMigrate(&layouts.Score{})
	stats.AutoMigrate(database.DB)

	osuapi.Enabled = conf.OsuApi.Enable
	osuapi.Key = conf.OsuApi.Key
}

func main() {
	web.Start(&config.Web)
	bancho.Start(&config.Bancho)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-exit

	log.Info("Exiting...")
}
