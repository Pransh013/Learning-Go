//  main package in Go is the entry point of the program, where the execution starts
package main

// format package: used for formatted I/O operations
import (
    "fmt"
    "math/rand"
    "math"
)

// Main fn: entry point for GO program
func main() {
    fmt.Println("Hello World")
    fmt.Println("some random number", rand.Intn(10))
    fmt.Println("value of pi", math.Pi)
    // math.pi: private variable (starting with lowercase) not accessible outside math package
    // math.Pi: public variable (starting with Uppercase) accessible outside math package
}