package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func main() {
	stringToReverse := "Hello world"
	strlength := len(stringToReverse)
	stringToReturn := ""

	for i := strlength - 1; i >= 0; i-- {
		stringToReturn = stringToReturn + string(stringToReverse[i])
	}

	fmt.Println(stringToReturn)
}
