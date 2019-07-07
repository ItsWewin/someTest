package main

import "fmt"

func main() {
	list := new([]int)
	*list = append(*list, 0, 1, 2)
	fmt.Println(list)

	// list := make([]int, 0)
	// list = append(list, 1)
	// fmt.Println(list)
}
