package utils

import (
	"fmt"
	"os"
	"strings"
)

func TildeToHome(path string) (string, error) {
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
