package pwgen

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Hello() {
	fmt.Println("Hello, world!")
}

func LoadDictionary(path string) []string {
	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Make sure you have specified a valid path for dictionary word list")
		os.Exit(1)
	}

	return strings.Split(string(content), "\n")
}

func GenerateWords(dict []string, numOfWords int) []string {
	var words []string

	for range make([]int, numOfWords) {
		var pickWord string

		capitalize := rand.IntN(2)

		if capitalize == 1 {
			pickWord = cases.Title(language.English, cases.Compact).String(dict[rand.IntN(len(dict))])
		} else {
			pickWord = dict[rand.IntN(len(dict))]
		}

		words = append(words, pickWord)
	}

	return words
}

func GenerateDigits(numOfDigits int) string {
	maxNum := int(math.Pow10(numOfDigits))

	return fmt.Sprintf("%d", rand.IntN(maxNum))
}

func GeneratePassword(words []string, digits string) string {
	var password string

	for i := range words {
		password += words[i]
		password += "-"
	}

	password = strings.ReplaceAll(password, "'", "")

	return password + digits
}
