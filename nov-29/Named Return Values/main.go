/*
Implement a function divideWithRemainder that takes two integers and
returns their quotient and remainder using named return values.
*/
package main

import (
	"errors"
	"fmt"
)

func divideWithRemainder(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("Cannot divide by zero.")
	} else {
		return x / y, x % y, nil
	}
}

func main() {
	quotient, remainder, err := divideWithRemainder(13, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quotient, remainder)
	}
}
