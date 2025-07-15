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

	// Accept input from the user
	fmt.Println("Enter text:")
	text.Scan()

	// clean the input text
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	cleanedText := re.ReplaceAllString(text.Text(), "")

	// to make case-insensitive
	words := strings.ToLower(cleanedText)

	// check palindrome
	isPalindrome := true

	for i := 0; i<len(words)/2; i++ {
		if words[i] != words[len(words)-1-i]{
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Yes ", text.Text(), " is Palindrome")
	} else{
		fmt.Println("No ", text.Text(), " is not Palindrome")
	}
}