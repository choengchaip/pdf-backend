package config

import "os"

type IConfig interface {
	MongoConfig() IMongoConfig
	Secret() string
	BaseURL() string
}

type Config struct {
	secret  string
	baseURL string
}

func NewConfig() IConfig {
	return &Config{}
}

func (c *Config) MongoConfig() IMongoConfig {
	return NewMongoConfig()
}

func (c *Config) Secret() string {
	if c.secret == "" {
		c.secret = os.Getenv("SECRET")
	}

	return c.secret
}

func (c *Config) BaseURL() string {
	if c.baseURL == "" {
		c.baseURL = os.Getenv("BASE_URL")
	}

	return c.baseURL
}
