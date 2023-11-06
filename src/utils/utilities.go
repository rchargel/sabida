package utils

import (
	"log"
	"os"
)

func GetEnv(variable string) string {
	v, ok := os.LookupEnv(variable)
	if !ok {
		log.Panicf("Missing environment variable %s", variable)
	}
	return v
}

func GetEnvOrDefault(variable, defaultValue string) string {
	v, ok := os.LookupEnv(variable)
	if !ok {
		return defaultValue
	}
	return v
}
