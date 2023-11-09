package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("Enter three integers (seperated by space" +
		"or newline):")

	fmt.Scan(&num1, &num2, &num3)
	fmt.Print("Largest number: ")
	result := num1

	if num2 > result {
		result = num2
	}
	if num3 > result {
		result = num3
	}

	fmt.Println(result)
}
