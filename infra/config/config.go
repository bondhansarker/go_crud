package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DbConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Schema string
}

type AppConfig struct {
	Name  string
	Port  string
	Page  int64
	Limit int64
	Sort  string
}

type Config struct {
	Db  *DbConfig
	App *AppConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func Db() *DbConfig {
	return config.Db
}

func LoadConfig() {
	loadEnvironmentVariables()
	setDefaultConfig()
}

func loadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func setDefaultConfig() {
	fmt.Println(os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_SCHEMA"))

	config.Db = &DbConfig{
		Host:   os.Getenv("MYSQL_HOST"),
		Port:   os.Getenv("MYSQL_PORT"),
		User:   os.Getenv("MYSQL_USER"),
		Pass:   os.Getenv("MYSQL_PASS"),
		Schema: os.Getenv("MYSQL_SCHEMA"),
	}
	config.App = &AppConfig{
		Name: os.Getenv("APP_NAME"),
		Port: os.Getenv("APP_PORT"),
	}
}
