package config

import (
	"github.com/infernalfire72/flame/config/bancho"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/config/web"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/infernalfire72/flame/log"
)

const FileName = "./conf.flm"

type Config struct {
	Database database.Config
	Bancho   bancho.Config
	Web      web.Config
	OsuApi   OsuApiConfig
}

type OsuApiConfig struct {
	Enable bool
	Key    string
}

var (
	Web    web.Config
	Bancho bancho.Config
)

func Load() (*Config, error) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Config File not found. Generating one for you.")
			Create()
			os.Exit(1337)
		}
		log.Error(err)
		return nil, err
	}

	conf := &Config{}

	if _, err = toml.Decode(string(data), conf); err != nil {
		log.Error(err)
		return nil, err
	}

	Web = conf.Web
	Bancho = conf.Bancho

	defer log.Info("Config has been loaded")
	return conf, nil
}

func Create() {
	file, err := os.Create(FileName)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	c := Config{
		Database: database.Config{
			ConnectionString: "user:pass@tcp(127.0.0.1:3306)/flame?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Bancho: bancho.Config {
			Port: 5001,
		},
		Web: web.Config{
			Port:           5002,
			ScreenshotPath: "./data/screenshots/%s.png",
		},
		OsuApi: OsuApiConfig{
			Key: "idk",
		},
	}

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		log.Error(err)
		return
	}

}
