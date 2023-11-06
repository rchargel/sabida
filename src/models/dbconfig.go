package models

import (
	"fmt"
	"strconv"

	"github.com/rchargel/sabida/utils"
)

type DbConfig struct {
	Username string
	Password []byte
	Scheme   string
	Host     string
	Port     uint16
}

func CreateDbConfig() *DbConfig {
	username := utils.GetEnv("DB_USERNAME")
	password := utils.GetEnv("DB_PASSWORD")
	scheme := utils.GetEnv("DB_SCHEME")
	host := utils.GetEnvOrDefault("DB_HOST", "localhost")
	port, _ := strconv.ParseUint(utils.GetEnvOrDefault("DB_PORT", "5432"), 10, 16)

	return &DbConfig{
		username,
		[]byte(password),
		scheme,
		host,
		uint16(port),
	}
}

func (dbConfig *DbConfig) GetConnectionStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.Username,
		string(dbConfig.Password),
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Scheme,
	)
}
