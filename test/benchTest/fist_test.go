//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 11:42 AM
//Copyright 2021 www.ucloud.cn

package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

/*
	golang 中：
		benchMark 压力测试单元
		形如：BenchmarkXxxx(t testing.B)
*/

func BenchmarkHello(b *testing.B) {
	for i:=0; i<b.N; i++ {
		fmt.Println("bench mark testing")
	}
}

func BenchmarkTmplExcute(b *testing.B) {
	b.ReportAllocs()
	templ := template.Must(template.New("test").Parse("hello, {{ . }}!"))
	b.RunParallel(func (pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "world")
		}
	})
}