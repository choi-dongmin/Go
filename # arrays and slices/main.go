package main

import "fmt"

func main() {
	names := []string{"DM", "BR", "BM"}
	names = append(names, "GN")
	fmt.Println(names)
}
