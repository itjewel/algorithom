package main
import "fmt"


func candy(ratings []int) int {
    n := len(ratings);
	candies := make([]int, n)
	for i:=0; i<n; i++{
		candies[i] = 1
	}

	for i:=1; i<n; i++{
		if ratings[i] > ratings[i-1]{
			candies[i] = candies[i-1] + 1
		}
	}

	for i:=n-2; i >=0; i--{		
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1]{
			candies[i] = candies[i+1] + 1
		}
	}

	total := 0
	for _, value := range candies {
		total += value
	}

	return total
	
}

func main() {
	ratings := []int{1,0,2}
	result := candy(ratings)
	fmt.Println(result)
}
