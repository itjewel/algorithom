/*
const s = "leetcode";
console.log(reverseVowels(s))
reverse only vowel
*/

package main

import (
	"fmt"
	"strings"
)

func reverseVowelsD(str string) string{
	var convertArray = strings.Split(str, "")
	var vowelArray = map[string]bool{"a": true,"A":true, "e": true,"E":true, "i":true,"I":true, "o":true,"O": true,"u": true, "U":true}
	left, right := 0, len(convertArray)-1
	for left<right {

		if !vowelArray[convertArray[left]]{
			left ++
			continue
		}

		if !vowelArray[convertArray[right]]{
			right --
			continue
		}
		convertArray[left], convertArray[right] = convertArray[right],convertArray[left];
		left++;
		right--
		
	}
	result:= strings.Join(convertArray, "")
	return result

}
func main(){

	str := "leetcode"
	result := reverseVowelsD(str);
	fmt.Println("Reverse Vowel", result)
}