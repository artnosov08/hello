package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println("hello world from Arthur", result)
}
// func to add integer
func add(x, y int) int {
	return x + y
}
