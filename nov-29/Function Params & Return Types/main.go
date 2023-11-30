/*
Write a function sumAndDifference that takes two integers and returns their sum and difference.
*/

package main

import "fmt"

func sumAndDifference(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	sum, diff := sumAndDifference(10, 5)
	fmt.Println(sum, diff)
}
