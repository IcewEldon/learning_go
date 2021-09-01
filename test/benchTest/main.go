//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 12:02 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func main() {
	benchmarkResult := testing.Benchmark(func (b *testing.B) {
		templ := template.Must(template.New("test").Parse("hello, {{.}}"))

		b.RunParallel(func (pb *testing.PB){
			var buf bytes.Buffer
			for pb.Next() {
				buf.Reset()
				templ.Execute(&buf, "world")
			}
		})
	})
	fmt.Printf("%s\t %s\n",
		benchmarkResult.String(), benchmarkResult.MemString())
}
