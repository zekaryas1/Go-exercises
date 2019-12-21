package main

import (
	"fmt"
	"strings"
)

func getFromDictionary(key string) string {
	preDefinedRoman := map[string]string{
		"1":    "I",
		"2":    "II",
		"3":    "III",
		"4":    "IV",
		"5":    "V",
		"6":    "VI",
		"7":    "VII",
		"8":    "VIII",
		"9":    "IX",
		"10":   "X",
		"20":   "XX",
		"30":   "XXX",
		"40":   "XL",
		"50":   "L",
		"60":   "LX",
		"70":   "LXX",
		"80":   "LXXX",
		"90":   "XC",
		"100":  "C",
		"200":  "CC",
		"300":  "CCC",
		"400":  "CD",
		"500":  "D",
		"600":  "DC",
		"700":  "DCC",
		"800":  "DCCC",
		"900":  "CM",
		"1000": "M",
	}

	return preDefinedRoman[key]
}

func stringToInt(stringNum string) int {
	result := 1
	switch stringNum {
	case "4":
		result = 4
	case "3":
		result = 3
	case "2":
		result = 2
	case "1":
		result = 1
	}

	return result
}

func Encode(word string) string {

	length := len(word) - 1
	result := ""
	for i := 0; i < len(word); i++ {
		if word[i] == '0' {
			length -= 1
			continue
		} else {
			if length < 1 {
				result = result + getFromDictionary(string(word[i]))
			} else {
				if length == 3 {
					current := strings.Repeat("M", stringToInt(string(word[i])))
					result = result + current
				} else {
					current := string(word[i]) + strings.Repeat("0", length)
					result = result + getFromDictionary(current)
				}
				length -= 1
			}
		}
	}
	return result
}

var toNumeric = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"L": 50, "C": 100, "D": 500, "M": 1000}

func decode(str string) (int, bool) {
	converted := 0
	lastVal := -1

	for _, i := range str {
		if _, ok := toNumeric[string(i)]; !ok {
			return 0, false
		}
		current := toNumeric[string(i)]

		if lastVal != -1 && lastVal < current && (current-lastVal)%4 != 0 && (current-lastVal)%9 != 0 {
			return 0, false
		}
		if lastVal != -1 && lastVal < current && ((current-lastVal)%4 == 0 || (current-lastVal)%9 == 0) {
			converted -= lastVal
			converted += current - lastVal
		} else {
			converted += current
		}
		lastVal = current
	}
	return converted, true
}

func main() {
	fmt.Println(Encode("1990") == "MCMXC")
	fmt.Println(Encode("94") == "XCIV")
	fmt.Println(Encode("230") == "CCXXX")
	fmt.Println(Encode("3450") == "MMMCDL")
	fmt.Println(Encode("1") == "I")
	fmt.Println(Encode("3999") == "MMMCMXCIX")
	fmt.Println(Encode("1200") == "MCC")

	fmt.Println(decode("MCMXC"))
}
