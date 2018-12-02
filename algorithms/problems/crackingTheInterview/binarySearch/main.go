package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lowerBound := 0
	upperBound := len(input)
	result := 0

	for lowerBound != upperBound {
		comparedValue := (lowerBound + upperBound) / 2

		if 9 == input[comparedValue] {
			result = comparedValue
			break
		} else if 9 < input[comparedValue] {
			upperBound = comparedValue
		} else {
			lowerBound = comparedValue + 1
		}
	}

	fmt.Println("Index:", result)
}
