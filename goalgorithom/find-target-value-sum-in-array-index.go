package main

import "fmt"

var arr = []int{1, 7, 2, 3, 2, 3, 6}
var target = 9
var createMap = make(map[int]int)

func TargetSumValue() {
	//var result = [][]int{}
	for i, value := range arr {
		var current = target - value
		if index, status := createMap[current]; status {
			fmt.Printf("get result %d %d", index, i)
			return
		}
		createMap[value] = i
		//if
	}
	// fmt.Printf("Loop %d", result)
	// fmt.Printf("Get Target vaue %d also get %d %d", target, arr, target)
}
func main() {
	TargetSumValue()
}
