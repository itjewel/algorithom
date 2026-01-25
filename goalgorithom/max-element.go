package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	majority := nums[0]
	max := make(map[int]int)
	for _, value := range nums {
		max[value]++
		if max[value] > max[majority] {
			majority = value

		}

	}
	//  fmt.Println(max)
	return majority

}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	// l1 := []int {2,7,11,15}
	// l2 := []int{7,0,8}
	result := majorityElement(nums)
	fmt.Println(result)
}
