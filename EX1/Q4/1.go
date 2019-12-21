package main

import "fmt"

func main() {
	listOfFloat64 := [] float64{12.5, 34.6, 56.89, 12.4, 3.5}
	slicedListOfFloat64 := listOfFloat64[0:4]

	length := float64(len(slicedListOfFloat64))
	var sum float64 = 0

	for _, value := range slicedListOfFloat64 {
		sum += value
	}

	fmt.Printf("The avarage is %f", sum/length)
}
