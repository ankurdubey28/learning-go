package main

import "fmt"

func main() {
	var p MyPrintable = 5
	fmt.Println(p)
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyPrintable int

func (m MyPrintable) String() string {
	return fmt.Sprintf("value is %d", m)
}
