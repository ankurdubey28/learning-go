package main

func main() {

}

type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T { // works
	return b.value
}

//func (b Box[T]) Map[V any](v V)T{   // gives error , uncomment it to understand.
//	return v
//}
