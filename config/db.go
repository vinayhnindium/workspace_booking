package config

import (
	"fmt"
	"os"
)

type dBConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var config = dBConfig{
	host:     "DB_HOST",
	port:     "DB_PORT",
	user:     "DB_USER",
	password: "DB_PASSWORD",
	dbname:   "DB_NAME",
}

func GetDBName() string {
	return os.Getenv(config.dbname)
}

func GetDBHost() string {
	return os.Getenv(config.host)
}

func GetDBPort() string {
	return os.Getenv(config.port)
}

func GetDBUser() string {
	return os.Getenv(config.user)
}

func GetDBPassword() string {
	return os.Getenv(config.password)
}

func GetDBConnectionURL() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetDBHost(),
		GetDBPort(),
		GetDBUser(),
		GetDBPassword(),
		GetDBName(),
	)
}
