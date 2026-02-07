package main
import "fmt"



func intToRoman(num int) string {
	symbleValue := map[int]int{
	0:1000,
	1:500,
	2:100,
	3:50,
	4:10,
	5:5,
	6:1,
}
symbleRoman := map[int]string{
	0: "M",
	1: "D",
	2: "C",
	3: "L",
	4: "X",
	5: "V",
	6: "I",
}
    output := ""
	for i:=0; i<len(symbleValue); i++{
		if symbleValue[i]>=num{
			output += symbleRoman[i]
			num -= symbleValue[i]
		}
	}
	// fmt.Println(len(num))
	return  output
}

func main() {
	nums := 3749
	result := intToRoman(nums)
	fmt.Println(result)
}
