package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	nums []int
	pos  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		pos:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.pos[val]; ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.pos[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.pos[val]
	if !ok {
		return false
	}

	last := this.nums[len(this.nums)-1]

	// swap with last
	this.nums[idx] = last
	this.pos[last] = idx

	// remove last
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.pos, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rs := Constructor()

	fmt.Println(rs.Insert(1))   // true
	fmt.Println(rs.Remove(2))   // false
	fmt.Println(rs.Insert(2))   // true
	fmt.Println(rs.GetRandom()) // 1 or 2
	fmt.Println(rs.Remove(1))   // true
	fmt.Println(rs.Insert(2))   // false
	fmt.Println(rs.GetRandom()) // 2
}
