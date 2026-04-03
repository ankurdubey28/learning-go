package main

import "fmt"

func main() {
	var a int = 20
	b := Convert[int, int](a) // specific case when type inference does not work
	fmt.Println(b)
}

// type element
type Integer interface {
	~int | ~int8
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}
