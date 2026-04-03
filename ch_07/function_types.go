package main

import "fmt"

func main() {
	var h Handler = HandlerFunc(printNum)
	h.Handle(42)
}

type Handler interface {
	Handle(int)
}

type HandlerFunc func(int)

func (h HandlerFunc) Handle(x int) {
	h(x)
}

func printNum(x int) {
	fmt.Println(x)
}
