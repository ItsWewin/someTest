package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println(slice1)
	fmt.Println(slice2)

	// slice can not used to '=='
	// if slice1 == slice2 {
	// 	fmt.Printf("")
	// }
}
