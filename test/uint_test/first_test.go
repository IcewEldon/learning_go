//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/30 1:54 PM
//Copyright 2021 www.ucloud.cn

package main

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func TestFib(t *testing.T) {
	var (
		in = 7
		expected = 13
	)

	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d, expected is %d", in, actual, expected)
	}
}
