package main
var array = []string{
	"apple",
	"apple",
	"banana",
	"banana",
	"apple",
	"date",
	"pomegranet",
	"pomegranet",
	"mango",
	"mangostring",
}
func DuplicateRemove() []string{
	seen := make(map[string]bool)
	var newArray []string;
	for _, item := range array{
		if !seen[item] {
			seen[item] = true;
			newArray = append(newArray, item)
		}
		// newArray = append(newArray, item)
		
	}
	
	return newArray;
	//fmt.Println(newArray)
		// fmt.Println("hello", item)
}