package main

import (
	"fmt"
	"math"
)

func main() {
	sl := []int{1, 2, 3, 5, 5, 6}
	sqr := Map(sl, func(t1 int) int {
		return t1 * t1
	})
	fmt.Println(sqr)
	filtered := Filter(sqr, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(filtered)
	val := Reduce(filtered, 0, func(i int, i2 int) int {
		return i + i2
	})
	fmt.Println(val)

	// Point2D comparison
	point1 := Pair[Point2D]{Point2D{X: 5, Y: 4}, Point2D{X: 4, Y: 7}}
	point2 := Pair[Point2D]{Point2D{X: 6, Y: 4}, Point2D{X: 2, Y: 9}}
	fmt.Println(FindCloser(point1, point2))
}

func Map[T1, T2 comparable](s []T1, f func(T1) T2) []T2 {
	sl := make([]T2, len(s))
	for i, v := range s {
		sl[i] = f(v)
	}
	return sl
}

func Filter[T comparable](s []T, f func(T) bool) []T {
	var sl []T
	for _, v := range s {
		if f(v) {
			sl = append(sl, v)
		}
	}
	return sl
}

func Reduce[T1 comparable](s []T1, initializer T1, f func(T1, T1) T1) T1 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

type Pair[T fmt.Stringer] struct {
	val1 T
	val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.val1.Diff(pair1.val2)
	d2 := pair2.val1.Diff(pair2.val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}

type Point2D struct {
	X, Y int
}

func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

func (p Point2D) String() string {
	return fmt.Sprintf("%d , %d", p.X, p.Y)
}
