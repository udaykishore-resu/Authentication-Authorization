package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string
	serverHost string
	serverPort int
}

func NewConfig() *Config {
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))

	return &Config{
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     dbPort,
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbName:     os.Getenv("DB_NAME"),
		serverHost: os.Getenv("SERVER_HOST"),
		serverPort: serverPort,
	}
}

func (c *Config) DatabaseConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.dbHost, c.dbPort, c.dbUser, c.dbPassword, c.dbName)
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.serverHost, c.serverPort)
}
