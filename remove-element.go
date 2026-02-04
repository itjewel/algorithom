package main
import(
	"fmt"
)

func removeElement(nums []int, val int) int {
    k := 0
    for _, v := range nums {
		if v != val {
			nums[k] = v 
			k++
		}
	}
	fmt.Println(nums)
	return k

}
	

func main(){
	nums := []int{3,2,2,3}
	val := 3
	removeElement(nums,val)
}