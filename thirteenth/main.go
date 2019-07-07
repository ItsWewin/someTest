package main

func GetValue() int {
	return 1
}

func main() {
	var i2 interface{}
	i := GetValue()
	i2 = i

	switch i2.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}
