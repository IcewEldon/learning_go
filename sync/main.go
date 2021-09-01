//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 3:36 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"learning_go/define"
	"learning_go/log"
	"sync"
	"time"
)

//lock

var (
	m *sync.RWMutex
	val = 0
)

func Read(i int) {
	zap.S().Info("function in read")
	m.RLock()
	defer m.RUnlock()

	time.Sleep(time.Second)
	zap.S().Info("val -> : ", val)
	time.Sleep(time.Second)
}

func Write(i int) {
	zap.S().Info("function in write")
	m.Lock()
	defer m.Unlock()

	val = 10
	time.Sleep(time.Second)
}

func Init() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)
}

func main() {
	Init()
	m = new(sync.RWMutex)
	go Read(1)
	go Write(2)
	go Read(3)
	time.Sleep(time.Second * 5)
}
