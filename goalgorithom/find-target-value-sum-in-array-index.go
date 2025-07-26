package main

import (
	"fmt"
)

var arr = []int{1, 7, 2, 3, 2, 3, 6}
var targetNum = 9

func targetValueIntex() {
	var makeMap = make(map[int]int)
	for i, num := range arr {
		var currentNum = targetNum - num
		if index, status := makeMap[currentNum]; status {
			fmt.Println("Hello", i, index)
			return
		} else {
			makeMap[num] = i
		}

	}

}
func main() {
	targetValueIntex()
}
