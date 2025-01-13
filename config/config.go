package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type KnownConfig struct {
	App struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"app"`
	Database struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Port     int    `json:"port"`
	} `json:"database"`
	Maildev struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"maildev"`
	Email struct {
		Username string `json:"username"`
		Password string `json:"password"`
		SmtpHost string `json:"smtpHost"`
		SmtpPort string `json:"smtpPort"`
	} `json:"email"`
}

func LoadConfig() (*KnownConfig, error) {
	var config KnownConfig

	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	return &config, nil
}
