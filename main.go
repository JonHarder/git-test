package main

import "fmt"

// / greet takes a name and prints a greeting.
func greet(name string) {
	fmt.Printf("Hello, %s\n", name)
	fmt.Printf("foo")
}

func inc(i int) int {
	return i + 3
}

// / This is the main method.
func main() {
	greet("Steve")
}
