package main

import "fmt"

// / greet takes a name and prints a greeting.
func greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}

// / This is the main method.
func main() {
	greet("Steve")
}
