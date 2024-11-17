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

var (
	searchCmd = &cobra.Command{
	Use: "search",
	Short: "Search functionality for VipinNotes",
	Run: nil,
 }

 searchByNameCmd = &cobra.Command{
	Use: "name",
	Short: " <name> Search user by name, it return all the users who's name contains name string",
	Run: searchByName,
 }

 searchByEmailCmd = &cobra.Command{
	Use: "email",
	Short: "<email> Search user by email, it return all the users who's email contains name email",
	Run: searchByEmail,
 }
 
 updateUserNameCmd = &cobra.Command{
	Use: "update",
	Short: "<email> Update name of user by email",
	Run: updateUserName,
 }

 getUserCmd = &cobra.Command{
	Use: "users",
	Short: "<n> Get last n users",
	Run: getUsers,
 }
)

func searchByName(cmd *cobra.Command, args []string)  {
	var name string
	fmt.Print("Enter Name : ")
	fmt.Scanln(&name)
	
	if name == "" {
		fmt.Println("Name should be there")
		return 
	}
	searchUserByName(name)
}

func searchByEmail(cmd *cobra.Command, args []string){
	var email string
	fmt.Print("Enter Email : ")
	fmt.Scanln(&email)

	if email == "" {
		fmt.Println("Name should be there")
		return 
	}

	searchUserByEmail(email)
}

func updateUserName(cmd *cobra.Command, args []string){
	var email, name string
	fmt.Print("Enter Email : ")
	fmt.Scanln(&email)
	fmt.Print("Enter Name : ")
	fmt.Scanln(&name)

	if email=="" || name==""{
		fmt.Println("All fields are requried")
		return
	}

	updateUserNameByEmail(name, email)

}

func getUsers(cmd *cobra.Command, args []string)   {
	var count int
	fmt.Print("Enter Count: ")
	_, err := fmt.Scanln(&count)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	if count <= 0 {
		fmt.Println("Count must be a positive number.")
		return
	}
	err= getUsersByCount(count)
	if err!=nil{
		fmt.Println(err.Error())
	}

}

func getAdminCredentials() map[string]interface{} {
	creds, err := utils.LoadCredentials()

		if err != nil {
			log.Fatalf("Failed to load credentials: %v", err)
		}

		return map[string]interface{}{
			"adminEmail": creds.Email,
			"adminPassword": creds.Password,
			"token": creds.AdminToken,
		}

}
func searchUserByName(name string) {
	apiUrl:= VipinNotesURL+"/admin/users/name"

	adminCredentials,err:=utils.LoadCredentials()

	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	
	reqBody,err:= json.Marshal(
			map[string]interface{}{
				"adminEmail": adminCredentials.Email,
				"adminPassword": adminCredentials.Password,
				"token": adminCredentials.AdminToken,
				"name":name,
			},
	)
	if err != nil {
		 fmt.Errorf("failed to encode credentials: %w", err)
		 return 
	}

	resp,err:= http.Post(apiUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		 fmt.Errorf("request failed: %w", err)
		 return 
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
			return 
		}
		
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error unmarshalling response body: %v", err)
		return 
	}

	if status, ok := response["status"].(string); ok && status == "error" {
		if message, exists := response["message"].(string); exists {
			fmt.Println(string(body))
			fmt.Printf("Error: %s\n", message)
		}
		return 
	} else {
		var formattedBody interface{}
		if err := json.Unmarshal(body, &formattedBody); err != nil {
			log.Fatal("Failed to unmarshal JSON:", err)
		}
		prettyJSON, err := json.MarshalIndent(formattedBody, "", "  ")
		if err != nil {
			log.Fatal("Failed to marshal JSON with indentation:", err)
		}
		fmt.Println(string(prettyJSON))
		return 
	}
}

func searchUserByEmail(email string) {
	apiUrl:= VipinNotesURL+"/admin/users/email"

	adminCredentials,err:= utils.LoadCredentials()

	if err!=nil{
		fmt.Println("Enable to load admin credentails, %v",err.Error())
		return
	}

	reqBody,err:= json.Marshal(map[string]interface{}{
			"adminEmail": adminCredentials.Email,
			"adminPassword": adminCredentials.Password,
			"token": adminCredentials.AdminToken,
			"email":email,
	})

	if err!=nil{
		fmt.Println("Enable to marshal request body, %v",err.Error())
		return
	}

	resp ,err := http.Post(apiUrl, "application/json", bytes.NewBuffer(reqBody))

	if err !=nil {
		fmt.Println("unable to fetch data, %v",err.Error())
		return
	}
  
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return 
	}

	var response map[string]interface{}

	if err = json.Unmarshal(body, &response); err!=nil{
		log.Fatalf("Error unmarshalling response body: %v", err)
		return
	}

	if status, ok := response["status"].(string) ; ok && status=="error"{
		if message,exists := response["message"].(string) ; exists{
			fmt.Println(string(body))
			fmt.Printf("Error: %s\n", message)
		}
	return 
	}else{
		var formattedBody map[string]interface{}

		if err := json.Unmarshal(body, &formattedBody); err != nil {
			log.Fatal("Failed to unmarshal JSON:", err)
		}
		prettyJSON, err := json.MarshalIndent(formattedBody, "", "  ")
		if err != nil {
			log.Fatal("Failed to marshal JSON with indentation:", err)
		}
		fmt.Println(string(prettyJSON))
		return 
	}
}

func updateUserNameByEmail(name, email string) {
	apiUrl := VipinNotesURL + "/admin/users/name"

	adminCredentials, err := utils.LoadCredentials()
	if err != nil {
		fmt.Println("Unable to load admin credentials")
		return
	}

	reqBody, err := json.Marshal(map[string]interface{}{
		"adminEmail":    adminCredentials.Email,
		"adminPassword": adminCredentials.Password,
		"token":         adminCredentials.AdminToken,
		"name":          name,
		"email":         email,
	})

	if err != nil {
		fmt.Println("Unable to marshal request body:", err)
		return
	}

	req, err := http.NewRequest("PUT", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error during request execution: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read response body: %v\n", err)
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error unmarshalling response body: %v", err)
		return
	}

	if status, ok := response["status"].(string); ok && status == "error" {
		if message, exists := response["message"]; exists {
			fmt.Println(string(body))
			fmt.Printf("Error: %s\n", message)
		}
		return
	} else {
		fmt.Println("Response:", response["result"])
	}
}

func  getUsersByCount(count int) error {
	apiURL := VipinNotesURL+"/admin/users"
	adminCredentials,err:=utils.LoadCredentials()

	reqBody, err := json.Marshal(map[string]interface{}{
		"adminEmail":    adminCredentials.Email,
		"adminPassword": adminCredentials.Password,
		"token":         adminCredentials.AdminToken,
		"number":        count, 
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
			return fmt.Errorf("Error reading response body: %v", err.Error())
		}

		var response map[string]interface{}
		if err := json.Unmarshal(body, &response); err != nil {
			log.Fatalf("Error unmarshalling response body: %v", err)
			return fmt.Errorf("Error unmarshalling response body: %v", err.Error())
		}

		if status, ok := response["status"].(string); ok && status == "error" {
			if message, exists := response["message"].(string); exists {
				fmt.Printf("Error: %s\n", message)
			}
			fmt.Println(string(body))
			return errors.New("You are not an admin")
		} else {
			var formattedBody interface{}
			if err := json.Unmarshal(body, &formattedBody); err != nil {
				log.Fatal("Failed to unmarshal JSON:", err)
			}
			prettyJSON, err := json.MarshalIndent(formattedBody, "", "  ")
			if err != nil {
				log.Fatal("Failed to marshal JSON with indentation:", err)
			}
			fmt.Println(string(prettyJSON))
			return nil
		}

}

func init(){
	RootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchByNameCmd)
	searchCmd.AddCommand(searchByEmailCmd)
	searchCmd.AddCommand(updateUserNameCmd)
	searchCmd.AddCommand(getUserCmd)
	
}
