package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	characters := "asSASA ddd dsjkdsjs dk"
	counts := make(map[string]int)
	var strength = len(characters)

	for i := 0; i < strength; i++ {
		var currentStr = string(characters[i])
		if currentStr != " " {
			counts[currentStr]++
		}
	}
	fmt.Printf("Number of byte in '%s' is %d \n", characters, utf8.RuneCountInString(characters))

	for key, value := range counts {
		fmt.Printf("%s appeared %d times \n", key, value)
	}
}
