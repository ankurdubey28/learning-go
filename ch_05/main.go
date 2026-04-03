package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(add(5, 6, 9))

	// function variables
	var myFuncVariable func(string) int
	myFuncVariable = f1
	fmt.Println(myFuncVariable("oneplus"))

	defer func() { fmt.Println(10) }()
	defer func() { fmt.Println(20) }()
	defer func() { fmt.Println(30) }()

	mp := map[int]int{
		1: 1,
	}
	sl := make([]int, 0, 3)
	modMap(mp)
	modSlice(sl)
	fmt.Println(mp)
	fmt.Println(sl)
}

func f1(a string) int {
	return len(a)
}

func add(val ...int) int {
	sum := 0
	for _, i := range val {
		sum += i
	}
	return sum
}

// named and blank returns
// when blank returns are executed , the last assigned values of return values are sent , which if not taken care of can be problematic
func divide(num, denom int) (res, remainder int, err error) { // these are named returns
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return // blank return
	}
	res, remainder = num/denom, num%denom
	return // blank return
}
