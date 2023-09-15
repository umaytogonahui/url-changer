package config

type Config struct {
	Env string `yaml:"env" evn:"ENV" envDefault:"development"`
}
