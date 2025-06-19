package cmd

import (
	"fmt"
	"os"

	"github.com/pwnderpants/go-dumb-pw/internal/pwgen"
	"github.com/spf13/cobra"
)

var numOfPasswords int
var numOfWords int
var numOfDigits int
var pathToDict string

var rootCmd = &cobra.Command{
	Use:   "go-dumb-pw",
	Short: "Generates a memorable dumb password",
	Long:  "Generate a memorable dumb password that is easy to type and remember",
	Run: func(cmd *cobra.Command, args []string) {
		for range make([]int, numOfPasswords) {
			wordList := pwgen.LoadDictionary(pathToDict)
			words := pwgen.GenerateWords(wordList, numOfWords)
			numPadding := pwgen.GenerateDigits(numOfDigits)

			fmt.Println(pwgen.GeneratePassword(words, numPadding))
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&numOfPasswords, "count", "c", 1, "Number of passwords to generate")
	rootCmd.Flags().IntVarP(&numOfWords, "words", "w", 2, "Number of words in the password")
	rootCmd.Flags().IntVarP(&numOfDigits, "digits", "d", 4, "Number of suffiex digits in the password")
	rootCmd.Flags().StringVarP(&pathToDict, "wordlist", "l", "/usr/share/dict/words", "Path to word list file")
}
