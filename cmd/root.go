package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "port",
	Short: "port - a cli to manage port",
	Long: "manage port",
	Run: func(cmd *cobra.Command, args []string)  {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr ,"There was an error while executing you CLI %s", err)
		os.Exit(1)
	}
}