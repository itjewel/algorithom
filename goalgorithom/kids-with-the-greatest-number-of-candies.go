package main

import "fmt"

var candies = [] int {2,3,5,1,3}
var extraCandies = 3

func LargeNumber() int{
	var max = candies[0]
	for _, candie := range candies {
		if candie > max {

			max = candie
		}

	}
	return max

}
func GreatestNumber() []bool{
 max := LargeNumber()
 var generateResult  []bool;
 for i:=0; i<len(candies); i++ {
		if candies[i]+extraCandies >= max {

			generateResult = append(generateResult, true);
		}else{
			generateResult = append(generateResult, false)
		}

	}
	return generateResult;
}

func main(){
	result:= GreatestNumber()
	fmt.Println("output", result)
}