package main

import "fmt"

var str1 = "abc"
var str2 = "pqr"

func mergerString() {
	maxNum := max(len(str1), len(str2))
	result := ""
	for i := 0; i < maxNum; i++ {
		if i < len(str1) {
			result += string(str1[i])
		}
		if i < len(str2) {
			result += string(str2[i])
		}
	}
	fmt.Println("", result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	mergerString()
}
