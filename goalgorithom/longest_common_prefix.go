package main
import "fmt"


func longestCommonPrefix(strs []string) string {
    if len(strs) <= 0{
		return ""
	}
	prefix := ""
	
	for i:=0; i<len(strs[0]); i++{
		ch := strs[0][i]
		for j:=1; j<len(strs); j++{
			if i >=len(strs[j]) || strs[j][i] != ch {
				return prefix
			}
		}
		prefix += string(ch)
		
	}

fmt.Println(prefix)
	return prefix
}

func main() {
	commonString := []string{"flower","flow","flight"}
	result := longestCommonPrefix(commonString)
	fmt.Println(result)
}
