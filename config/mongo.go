package config

import "os"

type IMongoConfig interface {
	Endpoint() string
	Port() string
	Username() string
	Password() string
}

type MongoConfig struct {
	endpoint string
	port     string
	username string
	password string
}

func NewMongoConfig() IMongoConfig {
	return &MongoConfig{}
}

func (m *MongoConfig) Endpoint() string {
	if m.endpoint == "" {
		m.endpoint = os.Getenv("MONGO_ENDPOINT")
	}
	return m.endpoint
}
func (m *MongoConfig) Port() string {
	if m.port == "" {
		m.port = os.Getenv("MONGO_PORT")
	}
	return m.port
}
func (m *MongoConfig) Username() string {
	if m.username == "" {
		m.username = os.Getenv("MONGO_USERNAME")
	}
	return m.username
}
func (m *MongoConfig) Password() string {
	if m.password == "" {
		m.password = os.Getenv("MONGO_PASSWORD")
	}
	return m.password
}
