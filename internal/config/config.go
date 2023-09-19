package config

import "time"

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer 			`yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"idle_timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() {
	configPath := os.Getenv(key:"CONFIG_PATH")
	if configPath == "" {
        log.Fatal("CONFIG_PATH is not set")
    }
}

// check if file exists
if _, err := os.Stat(configPath); os.IsNotExist(err) {
	log.Fatal(format: "config file %s does not exist", configPath)
}