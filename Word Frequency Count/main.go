package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	text := bufio.NewScanner(os.Stdin)
	freq := make(map[string]int)

	// Accept input from the user
	fmt.Println("Enter text:")
	text .Scan()

	// clean the input text
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	cleanedText := re.ReplaceAllString(text.Text(), "")

	// to make case-insensitive
	words := strings.Fields(strings.ToLower(cleanedText))

	// Count the frequency of each word
	for _, word := range words {
		freq[word]++
	}

	// Print the word frequency count
	fmt.Printf("Word\t\tFrequency\n")
	for word, count := range freq {
		fmt.Printf("%s\t\t%d\n", word, count)
	}
}