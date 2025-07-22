/* short variable declaration in if statement */
package main

import "fmt"
var people = map[string]int{
    "Alice":   25,
    "Bob":     30,
    "Charlie": 40,
}


func Practice(){
if age, found := people["charlie"]; found {
	fmt.Println("got it", age)
}

}