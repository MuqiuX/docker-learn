package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "dkr",
	Short: "my docker",
	Long: "description",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error")
	}
}