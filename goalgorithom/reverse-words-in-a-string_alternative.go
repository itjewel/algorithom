package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str) // Split by any whitespace
	left, right := 0, len(words)-1

	// In-place reverse the words slice
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ") // Join with single space
}

func main() {
	str := "The sky is blue"
	result := reverseWords(str)
	fmt.Println("Reversed:", result)
}
