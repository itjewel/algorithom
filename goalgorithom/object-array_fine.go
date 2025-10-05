package main

import (
	"encoding/json"
	"fmt"
)

type School struct {
	Name string `json:"name"`
	Age int `json:"age`
}
func CreateArray () []School{
	personaInfo := []School{
		{Name: "Jewel", Age: 15},
		{Name: "Kamal", Age: 25},
		
	}
	return personaInfo;
}
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	result := CreateArray()
	// Imagine this is a big array of objects
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "", Age: 30},
		{Name: "Bob", Age: 40},
		{Name: "", Age: 50},
		{Name: "Charlie", Age: 35},
	}

	// Filter only people with a Name
	filtered := filterWithName(people)

	// Convert to JSON to see the result
	jsonData, _ := json.MarshalIndent(filtered, "", "  ")
	fmt.Println(string(jsonData))
}

func filterWithName(people []Person) []Person {
	filtered := make([]Person, 0, len(people)) // preallocate for efficiency
	for _, p := range people {
		if p.Name != "" { // Only keep if Name has value
			filtered = append(filtered, p)
		}
	}
	return filtered
}
