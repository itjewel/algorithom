package main

import (
	"fmt"
	"strings"
)

func ReverseString(str string) string {
    arrayConveT := strings.Fields(str)
    left, right := 0, len(arrayConveT)-1;

    for left<right {
        arrayConveT[left], arrayConveT[right] = arrayConveT[right], arrayConveT[left]
        left++
        right--
    }
    
    return strings.Join(arrayConveT," ")

}

func main(){
    str := "This is a jewel farazi"
    result := ReverseString(str)
    fmt.Println("Reverse", result)
}