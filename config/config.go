package config

import (
	"github.com/bamzi/jobrunner"
	"sync"
)

var (
	global *Config
	once sync.Once
	Job jobrunner.Func
)

func NewConfig() *Config {
	once.Do(func() {
		global = &Config{}
	})
	return global
}
type Config struct {
	RunMode string
	HTTP HTTP
	Log Log
}

type HTTP struct {
	Host string
	Port int
	ShutdownTimeout int
}

type Log struct {
	 LogPath string
	 LogName string
	 Level int
	 Format string
}
