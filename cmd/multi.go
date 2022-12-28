package cmd

import (
	"github.com/cothi/port/port"
	"github.com/spf13/cobra"
)

var multi = &cobra.Command{
	Use:  "multi",
	Short: "multi",
	Run: func(cmd *cobra.Command, args []string) {
		port.KillPorts()
	},
}


