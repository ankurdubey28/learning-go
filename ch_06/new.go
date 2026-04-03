package main

import "fmt"

func main() {

	p1 := &Employee{}
	p2 := new(Employee)

	// both above syntaxes do the same thing
	fmt.Println(p1 == p2)   // false
	fmt.Println(*p1 == *p2) // true

	fmt.Println(p1)
	fmt.Println(p2)

	// how new behaves with non reference and reference types
	i := new(int)
	*i = 10 // perfectly fine, because there is nothing underlying to initialize before usage

	//var x int   // same as doing j:=new(int)
	//j:=&x

	m := new(map[string]int)
	(*m)["a"] = 2 // will give runtime panic because underlying structure is set to its zero value and not initialized

}

type Employee struct {
	Name   string
	Age    int
	Gender string
}
