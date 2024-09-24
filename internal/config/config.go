package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ptdewey/oolong/internal/utils"
)

var configFiles []string = []string{"oolong.json", ".oolong.json", ".oolongrc"}

// CHANGE: to list with template selection for create to allow multiple options
// - possibly allow multiple template directories?
type Config struct {
	NoteSources []string `json:"noteSources"`
	DefaultExt  string   `json:"defaultExt"`
	UseTemplate bool     `json:"useTemplate"`
	TemplateDir string   `json:"templateDir"`
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
		temp, err := utils.TildeToHome(src)
		if err != nil {
			// CHANGE: maybe switch to logging instead of print?
			fmt.Println(err)
			continue
		}
		cfg.NoteSources[i] = temp
	}

	cfg.TemplateDir, err = utils.TildeToHome(cfg.TemplateDir)
	if err != nil {
		cfg.TemplateDir = ""
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
			DefaultExt:  ".md",
		}
	}

	panic("Error: No config file specified and no fallback was found.")
}
