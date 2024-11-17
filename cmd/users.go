package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.co/vipinnotes-cli/utils"
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Fetch users for vipinnotes",
	Run: func(cmd *cobra.Command, args []string) {
		creds, err := utils.LoadCredentials()
		if err != nil {
			log.Fatalf("Failed to load credentials: %v", err)
		}

		resp, err := fetchUsers(creds)
		if err != nil {
			log.Fatalf("Failed to fetch users: %v", err)
		}

		fmt.Println("Users fetched successfully:")
		fmt.Println(resp)
	},
}

func init() {
	RootCmd.AddCommand(usersCmd)
}

// fetchUsers sends an authenticated request to fetch users
func fetchUsers(creds utils.Credentials) (string, error) {
	apiURL := "https://vipinnotes.com/api/admin/users" // Replace with your API endpoint

	reqBody := map[string]string{
		"email":       creds.Email,
		"password":    creds.Password,
		"admin_token": creds.AdminToken,
	}

	fmt.Println(reqBody)
	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned status: %s", resp.Status)
	}

	responseBody := new(bytes.Buffer)
	responseBody.ReadFrom(resp.Body)
	return responseBody.String(), nil
}
