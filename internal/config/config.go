package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type DBConf struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
}

type Config struct {
	DB DBConf
}

func Load(log *slog.Logger) (Config, error) {
	c := Config{}

	var err error

	c.DB, err = loadDbConf(log)
	if err != nil {
		return c, fmt.Errorf("error on load db config: %w", err)
	}

	return c, nil
}

func loadDbConf(log *slog.Logger) (DBConf, error) {
	dbConf := DBConf{}

	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) == 0 {
		log.Error("env DB_HOST is empty. Couldn't initialize application")
		return dbConf, errors.New("env DB_HOST is empty")
	}
	dbConf.Host = dbHost

	dbStrPort := os.Getenv("DB_PORT")
	if len(dbStrPort) == 0 {
		log.Error("env DB_PORT is empty. Use default value")
		dbStrPort = "5432"
	}

	dbPort, err := strconv.ParseInt(dbStrPort, 10, 64)
	if err != nil {
		log.Error("Error on parse config db port", err)
	}
	dbConf.Port = uint16(dbPort)

	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		log.Error("env DB_USER is empty. Couldn't initialize application")
		return dbConf, errors.New("env DB_USER is empty")
	}
	dbConf.User = dbUser

	dbPass := os.Getenv("DB_PASS")
	if len(dbPass) == 0 {
		log.Error("env DB_PASS is empty. Couldn't initialize application")
		return dbConf, errors.New("env DB_PASS is empty")
	}
	dbConf.Password = dbPass

	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		log.Error("env DB_NAME is empty. Couldn't initialize application")
		return dbConf, errors.New("env DB_NAME is empty")
	}
	dbConf.DBName = dbName

	return dbConf, nil
}
