package main

import "fmt"

type ME struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "mandu"}
	DM := ME{name: "DM", favFood: favFood, age: 21}
	fmt.Println(DM)
}
