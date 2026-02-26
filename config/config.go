package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Load() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	cfg := &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "go-starter-kit"),
			Port: getEnv("APP_PORT", "8080"),
			Env:  getEnv("APP_ENV", "development"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "go_starter_kit"),
		},
	}

	return cfg, nil
}

func (db *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
