package config

type BanchoConfig struct {
	Port int
}

type WebConfig struct {
	Port        int
	AllowedMods map[string]bool
}
