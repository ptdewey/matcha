package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var configFiles []string = []string{"notemgr.json", ".notemgr.json", ".notemgrrc"}

type Config struct {
	NoteSources []string `json:"noteSources"`
}

func ParseConfig() Config {
	configPath, err := findConfigurationFile()
	if err != nil {
		return defaultConfig()
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return defaultConfig()
	}

	var cfg Config
	if err = json.Unmarshal(contents, &cfg); err != nil {
		return defaultConfig()
	}

	if len(cfg.NoteSources) == 0 {
		return defaultConfig()
	}

	for i, src := range cfg.NoteSources {
		temp, err := tildeToHome(src)
		if err != nil {
			// CHANGE: maybe switch to log instead of print
			fmt.Println(err)
			continue
		}
		cfg.NoteSources[i] = temp
	}

	return cfg
}

func findConfigurationFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return "", err
	}

	for _, f := range configFiles {
		configPath := filepath.Join(home, f)
		if checkFileExists(configPath) {
			return configPath, nil
		}
	}

	return "", errors.New("No configuration file found.")
}

func checkFileExists(configPath string) bool {
	if info, err := os.Stat(configPath); err == nil {
		return !info.IsDir()
	}
	return false
}

func defaultConfig() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("Error getting user home directory: %s\n", err))
	}

	for _, cfgLocation := range configFiles {
		dir := filepath.Join(home, cfgLocation)
		if _, err := os.Stat(dir); err != nil {
			continue
		}

		return Config{
			NoteSources: []string{dir},
		}
	}

	// TODO: rename with better name
	panic("Error: No Note Manager config directory specified and no fallback was found.")
}

func tildeToHome(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return "", err
	}

	return strings.Replace(path, "~", home, 1), nil
}
