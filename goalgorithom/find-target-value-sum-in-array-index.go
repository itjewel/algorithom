package main

import "fmt"

var arr = []int{1, 7, 2, 3, 2, 3, 6}
var targetNum = 9

func findTargetValue() {
	//fmt.Println("Array:", arr)
	// HashMap approach: value -> index
	numMap := make(map[int]int)

	for i, num := range arr {
		complement := targetNum - num
		if idx, found := numMap[complement]; found {
			fmt.Printf("Target %d found at indices: [%d, %d]\n", targetNum, idx, i)
			return
		}
		numMap[num] = i
	}

	fmt.Println("No two numbers found with the target sum.")
}

func main() {
	findTargetValue()
}
