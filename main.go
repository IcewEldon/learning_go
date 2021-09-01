//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/25 5:35 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Func1() {

	go func(){
		for{
			Func1()
		}
	}()
}

func Func2() {
	fmt.Println("do func 2")
}

func main() {
	Func1()
	http.ListenAndServe("127.0.0.1:6060", nil)
}

