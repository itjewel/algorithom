package main
import "fmt"


func lengthOfLastWord(s string) int {
	n := len(s)
    result := 0
	for i:=n-1; i >=0; i--{
		if s[i] == ' '{
			if result >0 {
				break
			}
		}else{
			result++
		}
	}
	return  result
}

func main() {
	s := "Hello World "
	result := lengthOfLastWord(s)
	fmt.Println(result)
}
