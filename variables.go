package main

import "fmt"

func main() {
	var message string = "Hello, Go!"
	var age int = 23
	var isReady = true
	
	fmt.Println(message, age, isReady)
	
	// Short variable declaration (type inference)
	name := "Alice"
	count := 10
	
	fmt.Println(name, count)
	
	// Constants are values that cannot be changed after they are defined.
	const pi = 3.14165
	const greeting = "welcome!"
	
	fmt.Println(pi, greeting)
}