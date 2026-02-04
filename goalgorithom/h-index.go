package main
import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
    // sorting and counting sort
   sort.Sort(sort.Reverse(sort.IntSlice(citations)))
   h :=0
   for i :=0; i<len(citations); i++{
		if citations[i] >= i+1{
			h = i+1
		}else{
			break
		}
   }
	return h
}

func main(){
    l1 := []int {3,0,6,1,5}
    
   result := hIndex(l1)
    fmt.Println(result)
}