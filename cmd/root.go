package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pwnderpants/go-dumb-pw/internal/pwgen"
	"github.com/spf13/cobra"
)

var numOfPasswords int
var numOfWords int
var numOfDigits int
var pathToDict string
var template string

var rootCmd = &cobra.Command{
	Use:   "go-dumb-pw",
	Short: "Generates a memorable dumb password",
	Long:  "Generate a memorable dumb password that is easy to type and remember",
	Run: func(cmd *cobra.Command, args []string) {
		for range make([]int, numOfPasswords) {
			wordList := pwgen.LoadDictionary(pathToDict)
			words := pwgen.GenerateWords(wordList, numOfWords)
			numPadding := pwgen.GenerateDigits(numOfDigits)

			tpl := template
			if !isValidTemplate(template, numOfWords) {
				fmt.Fprintf(os.Stderr, "Invalid template passed in, using default format: {words}-{digits}\n")
				tpl = "{words}-{digits}"
			}

			fmt.Println(pwgen.GeneratePassword(words, numPadding, tpl))
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
	rootCmd.Flags().StringVarP(&template, "template", "t", "{words}-{digits}", "Template for password output. Examples: '{words}-{digits}', '{digits}-{words}', '{w1}-{w2}-{digits}'")
}

func isValidTemplate(template string, numWords int) bool {

	valid := regexp.MustCompile(`\{(words|digits|w[1-9][0-9]*)\}`)
	all := regexp.MustCompile(`\{[^}]+\}`)
	allPlaceholders := all.FindAllString(template, -1)

	for _, ph := range allPlaceholders {
		if !valid.MatchString(ph) {
			return false
		}

		if strings.HasPrefix(ph, "{w") && ph != "{words}" {
			n, err := strconv.Atoi(ph[2 : len(ph)-1])
			if err != nil || n < 1 || n > numWords {
				return false
			}
		}
	}

	return true
}
