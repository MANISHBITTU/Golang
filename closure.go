package main

import "fmt"

func main() {
	var fname, lname string = "Manish", "Sharma"
	fullname := func() string {
		return fname + " " + lname
	}
	fmt.Println(fullname())
}
