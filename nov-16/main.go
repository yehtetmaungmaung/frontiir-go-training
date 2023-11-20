package main

import "fmt"

func main() {
	var classes [][]int
	classA := []int{56, 65, 70, 59, 87, 28, 65, 39, 75, 60}
	classB := []int{52, 90, 100, 30, 60}
	classC := []int{90, 60, 80, 40, 90, 100, 30, 20, 0, 40, 50, 60, 97, 50, 44}

	classes = append(classes, classA, classB, classC)

	for i := range classes {
		for _, mark := range classes[i] {
			if mark == 100 {
				fmt.Println("Full Mark!", mark)
			} else if mark <= 99 && mark >= 80 {
				fmt.Println("Flying Color", mark)
			} else if mark <= 79 && mark >= 40 {
				fmt.Println("Pass", mark)
			} else if mark <= 39 && mark >= 0 {
				fmt.Println("Fail", mark)
			} else {
				fmt.Println("Invalid")
			}
		}
	}
}
