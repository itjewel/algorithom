package main
import "fmt"



func intToRoman(num int) string {
	symbleValue := map[int]int{
	0:1000,
	1:900,
	2:500,
	3:400,
	4:100,
	5:90,
	6:50,
	7:40,
	8:10,
	9:9,
	10:5,
	11:4,
	12:1,
}
symbleRoman := map[int]string{
	0: "M",
	1: "CM",
	2: "D",
	3: "CD",
	4: "C",
	5: "XC",
	6: "L",
	7: "XL",
	8: "X",
	9: "IX",
	10: "V",
	11: "IV",
	12: "I",
}
    output := ""
	for i:=0; i<len(symbleValue); i++{
		for symbleValue[i] <= num {
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
