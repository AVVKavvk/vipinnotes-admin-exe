package cmd

import (
	"fmt"

	"github.co/vipinnotes-cli/utils"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use: "logout",
	Short: "Logout from vipinnotes as admin",
	Run: logout,
}

func logout(cmd *cobra.Command, args []string )  {
	var isLogout string
	fmt.Println("Are you sure (Y/N)")
	fmt.Scanln(&isLogout)

	if isLogout=="Y" || isLogout=="y" {
		err:= utils.DeleteCredentials()

		if err!=nil{
			fmt.Printf("Somethings went wrong, %v\n",err.Error())
		}else{
			fmt.Println("Successfully logout")
		}
	}
}

func init()  {
	RootCmd.AddCommand(logoutCmd)
}