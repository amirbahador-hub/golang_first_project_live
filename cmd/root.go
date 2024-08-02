package main

import (
	"digikala/apis"
	"digikala/migrations"
	"fmt"
	"github.com/spf13/cobra"
)




  func main() {
  
	var cmdMigrate = &cobra.Command{
	  Use:   "migrate [string to print]",
	  Short: "Print anything to the screen",
	  Long: `print is for printing anything back to the screen.
  For many years people have printed back to the screen.`,
	  Run: func(cmd *cobra.Command, args []string) {
		migrations.RunMigrations()
	  },
	}
  
	var cmdRun = &cobra.Command{
	  Use:   "run [string to echo]",
	  Short: "Echo anything to the screen",
	  Long: `echo is for echoing anything back.
  Echo works a lot like print, except it has a child command.`,
	  Run: func(cmd *cobra.Command, args []string) {
		migrations.Setup()
		router := apis.GetRouter()
		if err := router.Start(":8080"); err != nil {
			fmt.Println("Nooooooooooo")
		}
	  },
	}
  
	
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdRun, cmdMigrate)
	rootCmd.Execute()
  }