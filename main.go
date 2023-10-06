package main

import "fmt"

func main() {
	result := add(1, 2)
	result2 := sub(2, 1)
	fmt.Println("hello world from Arthur", result, result2)
}
// func to add integer
func add(x, y int) int {
	return x + y
}
// func to subtract integer
func sub(x, y int) int {
	return x - y
}
