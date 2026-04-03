package main

import "fmt"

func main() {
	// normal usage
	myAdder := Adder{18}
	fmt.Println(myAdder.AddTo(5))

	// method value - assign method to variable or pass as parameter
	f1 := myAdder.AddTo
	fmt.Println(f1(5))

	// method expressions - create function from type itself
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 10))
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}
