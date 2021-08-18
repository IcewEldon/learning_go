//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/18 7:15 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings : ", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted : ", s)
}
