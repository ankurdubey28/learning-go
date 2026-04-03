package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	var p interface{}
	p = person{name: "ankur", age: 25}
	fmt.Println(p.(person).name)
	fmt.Println(p.(person).age)
}

type person struct {
	name string
	age  int
}

func doThings(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("j is nil")
	case int:
		fmt.Println("j is type of int")
	case string:
		fmt.Println("j is type of string")
	default:
		fmt.Println(j)
	}
}
