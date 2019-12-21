package main

import (
	"fmt"
	"sort"
	"strings"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func AreAnagram(s1, s2 string) bool {
	var r1 ByRune = StringToRuneSlice(s1)
	var r2 ByRune = StringToRuneSlice(s2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}

func findAnagram(dictionary []string, word string) []string {
	result := make([]string, 0, len(dictionary)+1)
	for _, value := range dictionary {
		if AreAnagram(strings.ToUpper(value), strings.ToUpper(word)) == true {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	result := findAnagram([] string{"Radium came", "Radiua cmme"}, "Madam Curie")
	fmt.Println(result)
}
