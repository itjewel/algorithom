package main

import (
	"fmt"
)

var array = []string{
	"apple",
	"potato",
	"apple",
	"orange",
	"lemon",
	"apple",
	"nut",
	"lemon",
}

func main() {
	frequency := make(map[string]int)
	for i := 0; i < len(array); i++ {
		frequency[array[i]]++
	}
	// fmt.Println(array)
	fmt.Println(frequency)
}
