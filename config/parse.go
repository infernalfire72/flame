package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/infernalfire72/flame/log"
)

const FileName = "./conf.flm"

type Config struct {
	Database DatabaseConfig
	Bancho   BanchoConfig
	Web      WebConfig
	OsuApi   OsuApiConfig
}

type OsuApiConfig struct {
	Enable bool
	Key    string
}

var (
	Web *WebConfig
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

	Web = &conf.Web

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
		Database: DatabaseConfig{
			Username: "root",
			Database: "akatsuki",
		},
		Bancho: BanchoConfig{
			Port: 5001,
		},
		Web: WebConfig{
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
