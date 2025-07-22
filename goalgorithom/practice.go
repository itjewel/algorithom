/* short variable declaration in if statement */
people := map[string]int{
    "Alice":   25,
    "Bob":     30,
    "Charlie": 40,
}

if age, found := people["charlie"]; found {
	fmt.Println("got it", age)
}