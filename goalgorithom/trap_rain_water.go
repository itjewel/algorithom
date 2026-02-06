package main
import "fmt"


func trap(height []int) int {
	left :=0
	right := len(height)-1
	leftMax :=0
	rightMax :=0
	water := 0
	for left < right{
		if height[left] < height[right]{
			if height[left] >= leftMax{
				leftMax = height[left]
			}else{
				water += leftMax - height[left]
			}
			left ++
		}else{
			if height[right] >= rightMax {
				rightMax = height[right]
			}else{
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

func main() {
	traping := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	result := trap(traping)
	fmt.Println(result)
}
