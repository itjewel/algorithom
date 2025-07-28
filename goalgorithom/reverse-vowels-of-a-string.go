package main

import (
	"fmt"
	"strings"
)

func reverseVowelsAl(str string) string{
	vowels := "aeiouAEIOU"
	 check := strings.Contains(vowels,"b")
	fmt.Println(check)
	return str;

}

func main(){
	str := "leetcode"
	result := reverseVowelsAl(str);
	fmt.Println(result)
}