package main

import (
	"fmt"
)

func MaximumFind(candies[]int, target int )[] bool{
   
   max:= candies[0];
   for _, value := range candies {
        if value > max {
            max = value;
        }
   }

   result := make([]bool, len(candies));
   for i, v := range candies{
      result[i] = v+target >= max
   }

 return result
}

func main(){
    candies := []int{12,1,12};
    extraCandies := 10
     result := MaximumFind(candies,extraCandies)
    fmt.Println("hello", result)
}