package main

import (
	"fmt"
	"reflect"
)

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	// sn1 和 sn2 字段相同、字段顺序相同、取值相同
	// sn1 == sn2
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	// sn1 和 sn3 的字段顺序不同，使用 == 比较编译报错
	// 下面的代码，不过不注释，会报错
	// sn3 := struct {
	// 	name string
	// 	age  int
	// }{name: "qq", age: 11}

	// if sn1 == sn3 {
	// 	fmt.Println("sn1 == sn3")
	// }

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	// 因为 slice、map、function 是不可比较类型
	// 下面的代码，不注释编译阶段会报错
	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }

	// 这种情况下可以使用 reflect.DeepEqual 比较
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	} else {
		fmt.Println("sm1 != sm2")
	}
}
