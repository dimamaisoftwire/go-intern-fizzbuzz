package main

import "fmt"

// This is our main function, this executes by default when we run the main package.
func main() {

	fmt.Println("Hello, World!")
	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

	// Put your code here...
}
