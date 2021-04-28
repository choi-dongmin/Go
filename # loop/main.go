package main

import (
	"fmt"
)

// func supperAdd(numbers ...int) {
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }

// func supperAdd(numbers ...int) {
// 	for _, number := range numbers {
// 		fmt.Println(number)
// 	}
// }
func supperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number //total = total + number
	}
	return total
}
func main() {
	result := supperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
