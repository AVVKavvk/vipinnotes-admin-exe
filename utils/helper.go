package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const credentialsFile = ".vipinnotes_admin_creds.json"

type Credentials struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	AdminToken string `json:"admin_token"`
}

func SaveCredentials(creds Credentials) {
	configPath, _ := filepath.Abs(credentialsFile)
	file, _ := json.MarshalIndent(creds, "", "  ")

	if err := os.WriteFile(configPath, file, 0600); err != nil {
		fmt.Printf("Failed to save credentials: %v\n", err)
		os.Exit(1)
	}
}

func LoadCredentials() (Credentials, error) {
	configPath, _ := filepath.Abs(credentialsFile)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Credentials{}, fmt.Errorf("failed to read credentials file: %w", err)
	}

	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return Credentials{}, fmt.Errorf("failed to parse credentials file: %w", err)
	}

	return creds, nil
}
func DeleteCredentials() error {
	configPath, _ := filepath.Abs(credentialsFile)
	return os.Remove(configPath)
}
