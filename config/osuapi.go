package config

import (
	"github.com/infernalfire72/flame/osuapi"
)

type OsuApiConfig struct {
	Key string
}

var ApiClient *osuapi.Client
