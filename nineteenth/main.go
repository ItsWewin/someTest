package main

import "fmt"

func foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}

	fmt.Println("non-empty interface")
}

func main() {
	var x *int
	foo(x)

	foo(nil)
}
