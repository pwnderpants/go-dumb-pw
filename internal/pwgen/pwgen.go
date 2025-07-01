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
	minNum := int(math.Pow10(numOfDigits - 1))
	maxNum := int(math.Pow10(numOfDigits) - 1)

	return fmt.Sprintf("%d", rand.IntN(maxNum-minNum)+minNum)
}

func GeneratePassword(words []string, digits string, template string) string {
	password := template
	for i, word := range words {
		placeholder := fmt.Sprintf("{w%d}", i+1)
		password = strings.ReplaceAll(password, placeholder, word)
	}
	password = strings.ReplaceAll(password, "{words}", strings.Join(words, "-"))
	password = strings.ReplaceAll(password, "{digits}", digits)
	password = strings.ReplaceAll(password, "'", "")

	return password
}
