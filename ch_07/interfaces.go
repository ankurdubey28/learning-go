package main

import "fmt"

func main() {
	var i all = nil
	println(i) // at runtime it is like this
	//type all struct {
	//	_type *_type          // pointer to runtime type descriptor
	//	data  unsafe.Pointer // pointer to actual data
	//}

	var x all = 42
	print(x)
	// at runtime
	//_type → int
	//data  → &42   (boxed on stack or heap)

	var i1 DoubleInt = 10
	var i2 DoubleInt = 10

	var dis1 = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}

	doubleCompare(&i1, &i2)  // why DoubleInt variables are called using & , because type DoubleInt
	doubleCompare(&i1, dis1) // has implemented the Double interface , but the receiver is *DoubleInt , not DoubleInt
	doubleCompare(dis1, dis2)

}

type all interface{}

type Doubler interface {
	Double()
}

type DoubleInt int

// *DoubleInt has Double implementation
func (d *DoubleInt) Double() {
	*d = *d * 2
}

// DoubleSlice has Double implementation
type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] *= 2
	}
}

// this function signature means Double method may be called on d1 and d2 , without knowing their concrete types
// thus compiler must guarantee at runtime that whatever is passed to doubleCompare , has to have callable Double method.
func doubleCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}
