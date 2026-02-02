package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	fmt.Println(k)

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums) // Output: [5 6 7 1 2 3 4]
}
