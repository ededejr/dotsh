package cmd

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	scriptDirName  = ".dotsh"
	apiKeyFileName = ".api-key"
)

func makeScriptDir() (string, error) {
	// Create directory for scripts if it doesn't exist
	usr, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	scriptDir := path.Join(usr, scriptDirName)
	err = os.MkdirAll(scriptDir, 0700)
	if err != nil {
		return "", err
	}
	return scriptDir, nil
}

func getAPIKey() (string, error) {
	scriptDir, err := makeScriptDir()
	if err != nil {
		return "", err
	}
	apiKeyFile := filepath.Join(scriptDir, apiKeyFileName)
	keyBytes, err := os.ReadFile(apiKeyFile)
	if err != nil {
		return "", err
	}
	apiKey := strings.TrimSpace(string(keyBytes))
	return apiKey, nil
}

func dasherize(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
