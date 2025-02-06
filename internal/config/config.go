package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/ptdewey/matcha/internal/utils"
)

var configFiles []string = []string{"matcha.toml", ".matcha.toml", ".matcharc"}

// CHANGE: to list with template selection for create to allow multiple options
// - possibly allow multiple template directories?
type Config struct {
	NoteSources []string `toml:"noteSources"`
	DefaultExt  string   `toml:"defaultExt"`
	UseTemplate bool     `toml:"useTemplate"`
	TemplateDir string   `toml:"templateDir"`
	NoteExts    []string `toml:"noteExts"`
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
	if err = toml.Unmarshal(contents, &cfg); err != nil {
		return defaultConfig()
	}

	if len(cfg.NoteSources) == 0 {
		return defaultConfig()
	}

	for i, src := range cfg.NoteSources {
		temp, err := utils.TildeToHome(src)
		if err != nil {
			log.Println(err)
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
		log.Println("Error getting user home directory:", err)
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
			UseTemplate: false,
			TemplateDir: "",
		}
	}

	panic("Error: No config file specified and no fallback was found.")
}
