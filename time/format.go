//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/19 10:42 上午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Printf

	t := time.Now()
	print("%s\n", t)

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	print("%s\n",t1)
	//当前时间的字符串，2006-01-02
	// 15:04:05据说是golang的诞生时间，固定写法
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	print("%s\n", timeStr)
}
