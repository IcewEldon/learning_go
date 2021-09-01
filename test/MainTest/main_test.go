//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 2:16 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var db struct {
	Dns string
}

func TestMain(m *testing.M) {
	db.Dns = os.Getenv("DATABASE_DNS")
	if db.Dns == "" {
		db.Dns = "dasdasdsd"
	}

	flag.Parse()
	exitCode := m.Run()

	db.Dns = ""
	os.Exit(exitCode)
}

func TestDatabase(t *testing.T) {
	fmt.Println(db.Dns)
}
