package main

import "fmt"

// 0 1 2 3 4 5 6 //
//-----river----//
// 1 2 4 0 6 5 3 //
// Result : 1 2 4 6 //

func main() {
	input := []int{
		1, 2, 4, 0, 6, 5, 3,
		// 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	res := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		res[i] = 1
	}

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] < input[j] {
				if res[j] < (res[i] + 1) {
					res[j] = res[i] + 1
				}
			}
		}
	}

	fmt.Println(res)
}
