package main

import "fmt"

func main() {
	//a := 5
	//var b Age
	//var c age
	//b := a // does not give error
	//c := a // gives error

	p1 := Person{
		name: "Ankur",
		age:  22,
	}
	p1.Greet()
	p1.Birthday() // compiler internally does (&p).Birthday() that is why no errors here
	// as to why compiler does this is because Birthday() has pointer receiver
	// and normally would require *person , but since P1 is a named variable means
	// it is addressable , compiler does the trick for us. Try to do same thing
	// anonymous struct and compiler will not do its magic.
	p2 := &Person{
		name: "Oneplus",
		age:  22,
	}
	p2.Greet() // compiler internally does (*p).Greet that is why no error here
	// again here explanation is simple because pointer always has access
	// to value with * operator so this is straight up safe to do so.
	p2.Birthday()

	var P3 *Person // nil instance
	P3.Birthday()
}

type Age = int // type alias
type age int   // type definition

type Person struct {
	name string
	age  int
}

// value receiver
func (p Person) Greet() {
	fmt.Println("Hi, I am", p.name, "and I am", p.age)
}

// pointer receiver
func (p *Person) Birthday() {
	if p == nil { // code for nil instances.
		fmt.Println("nil instance found")
		return
	}
	p.age++
}
