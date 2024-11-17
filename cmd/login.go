package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.co/vipinnotes-cli/utils"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as an admin user and store credentials locally",
	Run: login,
}

func login(cmd *cobra.Command, args []string) {
	var email, password, adminToken string
	fmt.Print("Enter Email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)
	fmt.Print("Enter Admin Token: ")
	fmt.Scanln(&adminToken)


	creds := utils.Credentials{
		Email:      email,
		Password:   password,
		AdminToken: adminToken,
	}

	err := authenticateUser(creds)
	if err != nil {
		fmt.Printf("Failed to log in: %v\n", err)
		return
	}

	utils.SaveCredentials(creds)
	fmt.Println("Logged in and credentials saved successfully!")
}

func authenticateUser(creds utils.Credentials) error {
	apiURL := "https://noteswebsiteserver.onrender.com/admin/users"

	reqBody, err := json.Marshal(map[string]interface{}{
		"adminEmail":    creds.Email,
		"adminPassword": creds.Password,
		"token":         creds.AdminToken,
		"number":        1, 
	})
	if err != nil {
		return fmt.Errorf("failed to encode credentials: %w", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}
		fmt.Println(string(body))

		var response map[string]interface{}
		if err := json.Unmarshal(body, &response); err != nil {
			log.Fatalf("Error unmarshalling response body: %v", err)
		}

		if status, ok := response["status"].(string); ok && status == "error" {
			if message, exists := response["message"].(string); exists {
				fmt.Printf("Error: %s\n", message)
			}
			return errors.New("You are not an admin")
		} else {
			utils.SaveCredentials(creds)
			fmt.Println("Authentication successful!")
		}

	// if resp.StatusCode != http.StatusOK {
	// 	body, _ := io.ReadAll(resp.Body)
	// 	return fmt.Errorf("server error: %s", string(body))
	// }
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("failed to read response body: %w", err)
	// }

	// var formattedBody interface{}
	// if err := json.Unmarshal(body, &formattedBody); err != nil {
	// 	log.Fatal("Failed to unmarshal JSON:", err)
	// }

	// // Marshal the data with indentation to print it in a pretty format
	// prettyJSON, err := json.MarshalIndent(formattedBody, "", "  ")
	// if err != nil {
	// 	log.Fatal("Failed to marshal JSON with indentation:", err)
	// }

	// // Print the pretty-printed JSON
	// fmt.Println("\nFormatted Response Body:")
	// fmt.Println(string(prettyJSON))

	// fmt.Println("Authentication successful!")
	return nil
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
