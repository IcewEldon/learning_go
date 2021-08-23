//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/19 11:56 上午
//Copyright 2021 www.ucloud.cn

package main

import "testing"

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != 2 {
		t.Errorf("IntMin(2, -2) = %d, want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		
	}
}