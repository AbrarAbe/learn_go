package main

    import "fmt"

    // Function to add two numbers
    func add(x int, y int) int {
        return x + y
    }

    // Function with multiple return values
    func swap(x, y string) (string, string) {
        return y, x
    }

    func main() {
        result := add(5, 3)
        fmt.Println("5 + 3 =", result)

        a := "Hello"
        b := "World"
        b, a = swap(a, b)
        fmt.Println(a, b) // Output: World Hello
    }