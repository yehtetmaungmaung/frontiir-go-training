/*
Implement a function divideWithRemainder that takes two integers and
returns their quotient and remainder using named return values.
*/
package main

import (
	"errors"
	"fmt"
)

func divideWithRemainder(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		return 0, 0, errors.New("Cannot divide by zero.")
	} else {
		return x / y, x % y, nil
	}
}

func main() {
	quotient, remainder, err := divideWithRemainder(13, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quotient, remainder)
	}
}
