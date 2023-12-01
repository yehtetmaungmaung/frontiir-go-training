package main

import "fmt"

func changeSlice(s *[]string) {
	(*s)[5] = "changed"
	(*s)[7] = "changed"
	(*s)[9] = "changed"
}

func main() {
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	changeSlice(&strSlice)
	fmt.Println(strSlice)
}
