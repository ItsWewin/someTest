package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// func (ua *UserAges) Get(name string) int {
// 	ua.Lock()
// 	defer ua.Unlock()
// 	if age, ok := ua.ages[name]; ok {
// 		return age
// 	}
// 	return -1
// }

func main() {
	age := make(map[string]int)
	age["wewin"] = 26
	userAges := UserAges{ages: age}
	for i := 0; i < 100; i++ {
		fmt.Printf("xxxx")
		go userAges.Add("wewin1", 27)
		go userAges.Get("wewin1")
	}
}
