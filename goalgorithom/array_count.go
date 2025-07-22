package main

import "fmt"

var Array2 = []string{
	"apple", "potato", "apple", "orange",
	"lemon", "apple", "nut", "lemon",
}

// func main(){

// 	res := ArrayCount()
// 	fmt.Println("res ",res)
// }

func ArrayCount() map[string]int {
	frequency := make(map[string]int)
	for _, item := range Array2 {
		frequency[item]++
	}
	fmt.Println(frequency)
	return frequency
}
