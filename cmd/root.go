package main 

import (
	"digikala/logger"
	"digikala/apis"
	"digikala/migrations"
	"net/http"
	"github.com/spf13/cobra"
)


func run(){

	myslog := logger.GetLogger()
	http.HandleFunc("/", apis.GetRoot)
	http.HandleFunc("/create_product", apis.CreateProduct)


	myslog.Info("Start Listening to 8000")
	serverErr := http.ListenAndServe(":8000", nil)
	if serverErr != nil {
		myslog.Error("Something is Wrong")
	}
}


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
		run()
	  },
	}
  
	
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdRun, cmdMigrate)
	rootCmd.Execute()
  }