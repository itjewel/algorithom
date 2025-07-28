package main

import (
	"fmt"
	"strings"
)

func vowelCheck(ch string) bool{
	list:= "aeiouAEIOU";
	result:= strings.Contains(list,ch);
	return result;
}
func reverseVowelsAl(str string) string{
	// vowels := "aeiouAEIOU"
	convertArray := strings.Split(str,"");
	left, right := 0, len(convertArray) -1
	for _, value := range convertArray{
		res :=vowelCheck(value)
		if res  {
			
			fmt.Println(res)
		}
	}
	//  check := strings.Contains(vowels,"b")
	// fmt.Println(check)
	return str;

}

func main(){
	str := "leetcode"
	result := reverseVowelsAl(str);
	fmt.Println(result)
}