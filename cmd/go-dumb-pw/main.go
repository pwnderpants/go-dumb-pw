package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var numToGen int

var rootCmd = &cobra.Command{
	Use:   "go-dumb-pw",
	Short: "Simple password generator",
	Long:  "A simple password generator that uses words that are easy to remember and type.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("For usage, run go-dumb-pw --h")
	},
}

func init() {
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
