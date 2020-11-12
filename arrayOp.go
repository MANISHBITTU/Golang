package main

import "fmt"

func main() {
	var arr [5]int
	i := 0
	fmt.Println("Enter 5 integer values:")
	for i = 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i])
	}
	sum := 0
	i = 0
	for i = 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("the sum of values is", sum)

}
