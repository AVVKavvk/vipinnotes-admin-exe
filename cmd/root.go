package cmd

import "github.com/spf13/cobra"

var VipinNotesURL = "https://noteswebsiteserver.onrender.com"

var RootCmd = &cobra.Command{
	Use: "vipinnotes",
	Short: "Admin page for vipinnotes",
	Run: nil,
}
func init()  {
	RootCmd.DisableAutoGenTag = true
}