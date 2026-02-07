package main
import "fmt"



func romanToInt(s string) int {
	symbolValue := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,

	}
	
	total :=0
	for i:=0; i<len(s); i++{
		if i+1 <len(s) && symbolValue[s[i]] < symbolValue[s[i+1]]{
			total -= symbolValue[s[i]]
		}else{
			total += symbolValue[s[i]]
		}
	}
	// fmt.Println(s[1])
	// fmt.Println(symbolValue)
	return total
	
}

func main() {
	s := "LVIII"
	result := romanToInt(s)
	fmt.Println(result)
}
