package main

import "fmt"

var (
	// 短变量的命名只能在函数内使用
	// size     := 1024
	size = 1024
	// 变量的命名应该使用驼峰形式
	//max_size = size * 2
	maxSize = size * 2
)

func main() {
	fmt.Println(size, maxSize)

	for i := 0; i < 10; i++ {
	loop:
		println(i)
		goto loop
	}

}
