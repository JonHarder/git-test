package main

import "fmt"

// / greet takes a name and prints a greeting.
func greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func inc(i int) int {
	return i + 1
}

// / This is the main method.
func main() {
	greet("Steve")
}
