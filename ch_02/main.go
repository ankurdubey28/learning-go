package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := `oneplus`
	fmt.Println(reflect.TypeOf(name))
	i := 32
	fmt.Println(reflect.TypeOf(i))
	//var x = 5   // compile time error
	const y = 5 // not error think why.

	//x:=1
	//y:=5
	//x,z:=4,6 // completely fine
	//x,y:=22,33 // not fine
}
