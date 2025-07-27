/*
const arr = [1, 2, 5, 7, 8, 9, 10];
const findIndexOf = 2;
*/
package main

import "fmt"

func findIndex(arr[]int, target int) int{
	for i, value := range arr {
		if(target == value){
			return i
		}
	}
	return -1
}

func main(){
	arr:= []int {1,1,3,2,7,8,9,10}
	target := 2
	result := findIndex(arr, target)
	fmt.Println("hello", result)
}