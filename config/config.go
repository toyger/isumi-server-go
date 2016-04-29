package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Bind  string `envconfig:"bind" default:":8080"`

}

var params *Config

func Parse() *Config {
	appConf := &Config{}
	if err := envconfig.Process("isumi", appConf); err != nil {
		panic(err)
	}

	params = appConf

	return appConf
}

func Get() *Config {
	if params == nil {
		panic("not initialize config parameter")
	}
	c := *params
	return &c
}
