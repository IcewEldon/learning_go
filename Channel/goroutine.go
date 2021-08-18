//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/18 3:50 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"time"
)

func Func(from string){
	for i:=0; i<len(from);i++ {
		fmt.Printf("From : %c\n", from[i])
	}
}

func main() {
	Func("test for function")

	go Func("go Func")

	go func(s string) {
		fmt.Println(s)
	}("going go")

	time.Sleep(time.Second)
}
