package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	ServerPort    int    `json:"server_port"`
	PostgresURL   string `json:"postgres_url"`
	JwtSecretKey  string `json:"jwt_secret_key"`
	LogPath       string `json:"log_path"`
}

func NewConfig() *Config {
	var config *Config

	content, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(err)
	}

	return config
}

// func NewConfig() *Config {
// 	var config Config

// 	viper.SetConfigFile("../config.json")

// 	err := viper.ReadInConfig()

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%+v", config)

// 	return &config
// }
// НАПОМНИТЬ ПОЧЕМУ НЕ РАБОТАЕТ
