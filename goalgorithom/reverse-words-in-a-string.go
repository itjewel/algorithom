/*
const s = "the sky is blue";
need to reverse word in the string
*/
package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	 arrayConver := strings.Fields(str)
	 count := len(arrayConver);
	stringD := ""
	 for i:=count-1; i >=0; i-- {
stringD += arrayConver[i]+" "
	 }
	return strings.TrimSpace(stringD) 
}

func main(){
	str := "The sky is blue"
	result := reverseString(str)
	fmt.Println("Reversed", result)
}