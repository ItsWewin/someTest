package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		fmt.Printf("stu: %v\n", &stu)
		m[stu.Name] = &stu
	}

}

func main() {
	pase_student()
}
