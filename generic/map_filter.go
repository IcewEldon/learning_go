package main


import (
	"fmt"
	"crypto/rand"
	"math/big"
	"strings"
)

func mapFunc[T any, M any] (a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))

	for i, v := range a {
		n[i] = f(v)
	}
	return n
}

func filterFunc[T any] (a []T, f func(T) bool) []T {
	var n []T
	for _, v := range a {
		if f(v) {
			n = append(n, v)
		}
	}
	return n
}

func main() {

	v1 := filterFunc(
			mapFunc([]int{1,2,3,4,5,6},
				func(v int) int {
					return v * v
				}),
				func(v int) bool {
					return v < 40
				})
	fmt.Println(v1)


	vs := filterFunc(mapFunc([]string{"a", "b", "c", "d", "e"},
					func(v string) string {
						n, _ := rand.Int(rand.Reader, big.NewInt(5))
						i := int(n.Int64())+1

						return strings.Repeat(v, i)
						}),
						func(v string) bool {
							return len(v) > 3
						})
	fmt.Println(vs)
}
