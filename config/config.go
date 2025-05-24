package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Token string `json:"token"`
}

var configPath = filepath.Join(os.Getenv("HOME"), ".zyracli", "config.json")

func SaveConfig(cfg Config) error {
	os.MkdirAll(filepath.Dir(configPath), 0755)
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(cfg)
}

func LoadConfig() (Config, error) {
	var cfg Config
	f, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&cfg)
	return cfg, err
}