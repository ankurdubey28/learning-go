package main

import "fmt"

type Printer interface {
	Print()
}

type ConsolePrinter struct{}

func (c ConsolePrinter) Print() {
	fmt.Println("printing ...")
}

type Service struct {
	printer Printer
}

func (s Service) DoWork() {
	s.printer.Print()
}

func main() {
	s := Service{printer: ConsolePrinter{}}
	s.DoWork()
}
