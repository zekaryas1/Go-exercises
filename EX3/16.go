package main

import "fmt"

func helperReverse(s []byte) string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(s)
}

func reverseParentheses(s string) string {
	ss := []byte(s)
	brackets := []int{}

	for i := 0; i < len(ss); i++ {
		if ss[i] == byte('(') || ss[i] == byte(')') {
			brackets = append(brackets, i)
		}
	}
	start := 0
	end := 0
	for i := 0; i < len(brackets); i++ {
		if s[brackets[i]] == byte('(') && end == 0 {
			start = brackets[i]
		}
		if s[brackets[i]] == byte(')') {
			end = brackets[i]
			tend := ""
			tend = s[0:start] + helperReverse([]byte(s[start+1:end]))
			if end+1 < len(s) {
				tend += s[end+1:]
			}
			return reverseParentheses(tend)

		}
	}

	return s
}

func main() {
	fmt.Println(reverseParentheses("(bar)") == "rab")
	fmt.Println(reverseParentheses("foo(bar)baz") == "foorabbaz")
	fmt.Println(reverseParentheses("foo(bar(baz))blim") == "foobazrabblim")
}
