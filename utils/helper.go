package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Path to the credentials file
const credentialsFile = "vipinnotes_admin_creds.json"

// Credentials structure
type Credentials struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	AdminToken string `json:"admin_token"`
}

// SaveCredentials stores the user's credentials in a local file
func SaveCredentials(creds Credentials) {
	configPath, _ := filepath.Abs(credentialsFile)
	file, _ := json.MarshalIndent(creds, "", "  ")

	if err := os.WriteFile(configPath, file, 0600); err != nil {
		fmt.Printf("Failed to save credentials: %v\n", err)
		os.Exit(1)
	}
}

// LoadCredentials reads the credentials from the local file
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

// DeleteCredentials removes the credentials file (used for logout)
func DeleteCredentials() error {
	configPath, _ := filepath.Abs(credentialsFile)
	return os.Remove(configPath)
}
