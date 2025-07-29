/*
1) need a algorithom develop in a string to vowel reverse
logic
* string to array convert
* vowel check
* left to right vowel check
* right to left vowel check
* then reverse
* array to string covert
* return value

// str := "applekhabo"
*/

package main

import (
	"fmt"
	"strings"
)

// func vowCheck(ch string) bool {
// 	str := "AaEeIiOoUu"
// 	return strings.Contains(str, ch)
// }
func reve(str string) string {
	vowCheck := map[string] bool {
		"a": true,"A": true,"e": true,"E": true,"i": true,"I": true,"o": true,"O": true,"u": true,"U": true,
	}
	convertArr := strings.Split(str,"")
	left, right := 0, len(convertArr) - 1
	
	for left<right {
		
		if !vowCheck[convertArr[left]] {
			left++
			continue;
		}
		if !vowCheck[convertArr[right]]{
			right--
			continue;
		}
		convertArr[left],convertArr[right] = convertArr[right], convertArr[left]
		left++;
		right--
		
		// fmt.Println(rr)
		// vowelleft = strings.Contains(convertArr, convertArr[left])
		// vowelright = strings.Contains(convertArr, convertArr[right])
		// fmt.Println(vowelleft, vowelright)
	}
	return strings.Join(convertArr,"")
}


func main(){
str := "Applekhabo"
result := reve(str)
fmt.Println(str, "-> ", result)
}

