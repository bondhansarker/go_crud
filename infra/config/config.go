package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Schema string
}

type RedisConfig struct {
	Host              string
	Port              string
	Pass              string
	Db                int
	AccessUuidPrefix  string
	RefreshUuidPrefix string
	UserPrefix        string
	TokenPrefix       string
	Ttl               int // seconds
}

type JwtConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	ContextKey         string
}

type AppConfig struct {
	Name  string
	Port  string
	Page  int64
	Limit int64
	Sort  string
}

type Config struct {
	Db    *DbConfig
	App   *AppConfig
	Jwt   *JwtConfig
	Redis *RedisConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func Db() *DbConfig {
	return config.Db
}

func Redis() *RedisConfig {
	return config.Redis
}

func Jwt() *JwtConfig {
	return config.Jwt
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

	Ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))
	DbName, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	config.Redis = &RedisConfig{
		Host:              os.Getenv("REDIS_HOST"),
		Port:              os.Getenv("REDIS_PORT"),
		Db:                DbName,
		Pass:              os.Getenv("REDIS_PASSWORD"),
		AccessUuidPrefix:  os.Getenv("REDIS_ACCESS_UUID_PREFIX"),
		RefreshUuidPrefix: os.Getenv("REDIS_REFRESH_UUID_PREFIX"),
		UserPrefix:        os.Getenv("REDIS_USER_PREFIX"),
		TokenPrefix:       os.Getenv("REDIS_TOKEN_PREFIX"),
		Ttl:               Ttl,
	}

	config.Jwt = &JwtConfig{
		AccessTokenSecret:  "access_token_secret",
		RefreshTokenSecret: "refresh_token_secret",
		AccessTokenExpiry:  300,
		RefreshTokenExpiry: 10080,
		ContextKey:         "user",
	}
}
