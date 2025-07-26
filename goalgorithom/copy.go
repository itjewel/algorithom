package main

import "fmt"

var arre = []int{1, 7, 2, 3, 2, 3, 6}
var targetNume = 9

func FindTargetValue() {
	//fmt.Println("Array:", arre)
	// HashMap approach: value -> index
	numMap := make(map[int]int)

	for i, num := range arre {
		complement := targetNume - num
		if idx, found := numMap[complement]; found {
			fmt.Printf("Target %d found at indices: [%d, %d]\n", targetNume, idx, i)
			return
		}
		numMap[num] = i
	}

	fmt.Println("No two numbers found with the target sum.")
}

func Copy() {
	findTargetValue()
}
