//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/19 11:43 上午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.TempFile("", "sample")
	check(err)

	fmt.Println("temp file name : ", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("temp dir name : ", dname)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
