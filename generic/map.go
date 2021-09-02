package main

import (
	"fmt"
	)

func mapFunc[T any, M any] (a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))

	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func main() {
	v1 := []int{ 1, 2, 3, 4, 5}
	vs := mapFunc(v1, func(v int) int {
		return v*v
	})
	fmt.Println(vs)
}

