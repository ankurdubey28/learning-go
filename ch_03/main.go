package main

import "fmt"

func main() {

	var x [3]int // zero array
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
	// sparse array
	var arr = [10]int{1, 4: 5, 9: 11}
	fmt.Println(arr)
	// spread operator
	y := [...]int{5, 5, 5} // ... operator picks length from declaration itself
	fmt.Println(len(y))
	// comparison using == and !=
	fmt.Println(x == y) // for == to say true , both size and elements of both array should be same.
	//size := 9
	//var nums [size]int // invalid
	const size = 9
	var nums [size]int // completely valid
	fmt.Println(nums)

	var sl = []int{1, 2, 3, 4}
	fmt.Println(sl)

	var nl []int
	fmt.Println(nl)

	sl1 := []int{1, 2, 3}
	sl2 := append(sl1, sl...)
	fmt.Println(sl2)

	//sl2:=make([]int,2,1) // compile time error if len > cap

	a := []int{1, 2, 3, 4, 5}
	// full slice expression
	s := a[1:3:3]               // max = 3
	fmt.Println(len(s), cap(s)) // 2, 2

	s = append(s, 99)
	fmt.Println(a) // [1 2 3 4 5] ✅ safe // since the full slice expression limits the capacity of s, append cannot grow s within the original backing array and is therefore forced to allocate a new backing array
}
