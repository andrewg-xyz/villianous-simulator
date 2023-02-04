package cmd

import (
	"fmt"

	"github.com/andrewg-xyz/villianous-simulator/src/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "villianous [COMMAND]",
	Short: "villianous is a CLI tool for running villianous simulations",
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := cmd.Flags().GetString("villian")
		state := internal.CreateGame(data)
		fmt.Println("state: ", state)
	},
}

// Execute is the entrypoint for the CLI.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}


func init() {
	rootCmd.PersistentFlags().StringP("villian", "v", "", "villian file path")
}