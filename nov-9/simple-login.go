package main

import "fmt"

func main() {
	var username, password string

	fmt.Println("Enter user name and password " +
		"(separated by space or newline:)")

	fmt.Scan(&username, &password)
	if username == "admin" && password == "password123" {
		fmt.Println("Login successful.")
	} else {
		fmt.Println("Login failed.")
	}
}
