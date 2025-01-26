package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type BranchConfig struct {
	Prefix  string `json:"prefix"`
	Default string `json:"default"`
}

type CommitConfig struct {
	NoVerify bool `json:"noVerify"`
}

type Config struct {
	Branch BranchConfig `json:"branch"`
	Commit CommitConfig `json:"commit"`
}

func GetConfig() (Config, error) {
	defaultConfig := Config{
		Branch: BranchConfig{
			Prefix:  "",
			Default: "",
		},
		Commit: CommitConfig{
			NoVerify: false,
		},
	}

	configDir := getConfigDir()
	appConfigPath := filepath.Join(configDir, AppName, ConfigFileName)

	file, err := os.Open(appConfigPath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Config file not found. Using defaults.")
			return defaultConfig, nil
		}
		fmt.Println("Error opening config file:", err)
		return defaultConfig, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return defaultConfig, nil
	}

	return config, nil
}

func getConfigDir() string {
	if runtime.GOOS == "darwin" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding user home directory for config:", err)
			os.Exit(1)
		}
		return filepath.Join(homeDir, ".config")
	}

	configDir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println("Error finding user home directory for config:", err)
		os.Exit(1)
	}

	return configDir
}
