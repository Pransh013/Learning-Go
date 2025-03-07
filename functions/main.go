package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// functions in GO can return multiple values, it is done by defining multiple return types in ()
func addSubtract(x, y int) (int, int) { // If multiple parameters have same type, omit type for all but the last one
	sum := x + y
	diff := x - y
	return sum, diff
}

// Named Return
// func addSubtract(x, y int) (sum, diff int) {
// 	sum = x + y
// 	diff = x - y
// 	return          // No need to explicitly return variables
// }

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func main() {
	fmt.Println(add(2, 7))

	sum, diff := addSubtract(10, 5)
	fmt.Println(sum, diff)

	plus := add // Assign function to a variable
	fmt.Println(plus(2, 5))

	// Anonymous function assigned to a variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println(multiply(2, 10))

	// Higher Order Functions
	result := applyOperation(10, 20, add)
	fmt.Println(result) 
}
