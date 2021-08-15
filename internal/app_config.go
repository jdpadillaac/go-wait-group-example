package internal

import "os"

type AppConfig struct {
	Env           string
	Version       string
	BaseUrl       string
	MysqlEngine   string
	MySqlServer   string
	MySqlPort     string
	MySqlUser     string
	MySqlPassword string
	MySqlDataBase string
	AppPort       string
}

func SetAppConfig() *AppConfig {
	return &AppConfig{
		AppPort:       os.Getenv("APP_PORT"),
		Env:           os.Getenv("ENV"),
		Version:       os.Getenv("VERSION"),
		BaseUrl:       os.Getenv("BASE_URL"),
		MysqlEngine:   os.Getenv("MYSQL_ENGINE"),
		MySqlServer:   os.Getenv("MYSQL_SERVER"),
		MySqlPort:     os.Getenv("MYSQL_PORT"),
		MySqlUser:     os.Getenv("MYSQL_USER"),
		MySqlPassword: os.Getenv("MYSQL_PASSWORD"),
		MySqlDataBase: os.Getenv("MYSQL_DATABASE"),
	}
}
