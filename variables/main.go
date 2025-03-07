package main

import (
	"fmt"
)

// Declare without initialization (defaults to zero value), must be assigned inside a function
var name string

// Explicit Type with initialization
var age int = 20

// Type Inference
var name2 = "Jane"

// Constant, cannot be reassigned
const pi = 3.14

// Multiple variable declaration
var (
	fName           = "Alice"
	score   float64 = 95.5
	success         = true
)

func main() {
	name = "John" // Assign inside the function
	age2 := 25    // Short declaration, must be used inside a function

	age := 35 // Variable Shadowing, this hides the global age variable

	// Multiple variable declaration using shorthand
	country, city, id := "India", "Lucknow", 1

	fmt.Println("hello ", name, age, name2, age2, pi)
	fmt.Println("hello ", fName, score, success)
	fmt.Println(country, city, id)
}
