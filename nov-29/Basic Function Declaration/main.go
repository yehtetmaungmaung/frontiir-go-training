/*
Write a function named greet that takes a name as a parameter and prints a greeting message.
Declare a function calculateSquare that takes an integer as a parameter and returns its square.
*/

package main

import "fmt"

func main() {
	greet("hello")
	fmt.Println(calculateSquare(10))
}

func greet(msg string) {
	fmt.Println(msg)
}

func calculateSquare(num int) int {
	return num * num
}
