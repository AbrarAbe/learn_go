package main

import "fmt"

func main() {
	age := 23
	
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
	
	if num := 10; num > 5 {
		fmt.Println("num is grrater than 5")
	}
	
	// Classic for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // While-like loop
    sum := 0
    for sum < 10 {
        sum += 2
    }
    fmt.Println("Sum:", sum)

    // Infinite loop (use 'break' to exit)
    // for {
    //     fmt.Println("This will run forever (unless you break)...")
    //     break // Exit the loop
    // }
    
    // Looping through array/slice
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

     // Looping through a string (rune by rune)
    message := "Hello"
    for index, char := range message {
        fmt.Printf("Index: %d, Char: %c\n", index, char)
    }
}