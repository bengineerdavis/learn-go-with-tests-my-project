package main

import "fmt"

// domain code
// never see the "outside world"
func Hello() string {
	return "Hello, world"
}

// code with side effects
// main function must always be defined as the entrypoint
// to a program's logic
func main() {
	fmt.Println(Hello())
}
