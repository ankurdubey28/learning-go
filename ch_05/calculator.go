package main

import "fmt"

func main() {
	opsMp := map[string]calFunc{
		"+": sum,
		"-": sub,
		"*": mul,
		"%": div,
	}
	// say input has given i=5,j=6 and ops=sum
	fmt.Println(opsMp["+"](5, 6))
}

func sum(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

type calFunc func(int, int) int
