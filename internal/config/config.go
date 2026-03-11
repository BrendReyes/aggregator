package config

import (
	"encoding/json"
	"path/filepath"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error){
    configPath, err := getConfigFilePath()
    if err != nil {
        return Config{}, err
    }

    file, err := os.Open(configPath)
    if err != nil {
        return Config{}, err
    }
    defer file.Close()

    var config Config
    err = json.NewDecoder(file).Decode(&config) //read then put the data in the struct
    if err != nil {
        return Config{}, err
    }

    return config, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(homeDir, configFileName)

	return configPath, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
    if err != nil {
        return err
    }

    file, err := os.Create(configPath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(cfg)
    if err != nil {
        return err
    }

    return nil
}