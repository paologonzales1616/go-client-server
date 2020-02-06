package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//envConfig represents database server and credentials
type envConfig struct {
	Port string
}

//Env holds environment variables
var Env *envConfig

//ReadEnv and parse the .env file
func ReadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Env = &envConfig{
		Port: getEnv("PORT", "3000")}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
