package array_count

import "fmt"

var Array2 = []string{
	"apple", "potato", "apple", "orange",
	"lemon", "apple", "nut", "lemon",
}

func ArrayCount() {
	frequency := make(map[string]int)
	for _, item := range Array2 {
		frequency[item]++
	}
	fmt.Println(frequency)
}
