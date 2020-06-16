package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/infernalfire72/flame/log"
)

const FileName = "./conf.flm"

type Config struct {
	Bancho		BanchoConfig
	Database	DatabaseConfig
}

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
		Bancho:		BanchoConfig {
			Port:		5001,
		},
		Database:	DatabaseConfig {
			Username:	"root",
			Database:	"akatsuki",
		},
	}

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		log.Error(err)
		return
	}


}