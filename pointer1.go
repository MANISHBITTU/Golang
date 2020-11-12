package main

import "fmt"

func main() {
	var x int = 7
	var y *int
	y = &x
	*y = 9
	fmt.Println(x)

}
