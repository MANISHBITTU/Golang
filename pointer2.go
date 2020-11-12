package main

import "fmt"

func main() {
	var x int = 17
	var y int = 9

	fmt.Println(add(&x, &y))

}

func add(a *int, b *int) int {
	return *a + *b
}
