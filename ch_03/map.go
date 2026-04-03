package main

import (
	"fmt"
)

func main() {
	mp := map[int]int{}
	mp[1] = 1
	mp[2] = 2
	mp[3] = 3
	for i, k := range mp {
		fmt.Println(i, k)
	}
	val, ok := mp[88]
	if !ok {
		fmt.Println("nahi hai")
	} else {
		fmt.Println(val)
	}
	delete(mp, 1)
	fmt.Println(mp)
}
