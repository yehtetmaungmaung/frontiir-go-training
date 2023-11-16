package main

import "fmt"

func main() {
	var sum int

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum = sum + i
		} else {
			sum = sum - i
		}
	}

	fmt.Println("Sum:", sum)
}
