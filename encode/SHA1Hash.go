//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/19 11:11 上午
//Copyright 2021 www.ucloud.cn

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	/*
		创建hash encodes
		New()
		Write() : 字节序列
		Sum() : 接受一个参数，一般为nil
	*/

	strs := "this is a hash string test"

	hash := sha1.New()
	hash.Write([]byte(strs))

	encode := hash.Sum(nil)

	fmt.Println(encode)
}
