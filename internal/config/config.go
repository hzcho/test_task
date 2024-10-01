package config

import (
	"time"
)

type Config struct {
	DB     `yml:"db" env-required:"true"`
	Server `yml:"server" env-required:"true"`
}

type Server struct {
	Port      string        `yml:"port" env-required:"true"`
	ReadTime  time.Duration `yaml:"read_time" env-required:"true"`
	WriteTime time.Duration `yaml:"write_time" env-required:"true"`
}

type DB struct {
	Username string `yml:"username" env-required:"true"`
	Host     string `yml:"host" env-required:"true"`
	Port     string `yml:"port" env-required:"true"`
	DBName   string `yml:"dbname" env-required:"true"`
	SSLMode  string `yml:"sslmode" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}
