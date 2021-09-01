//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 5:29 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"learning_go/define"
	"learning_go/log"
	"sync"
	"time"
)

var (
	locker = new(sync.Mutex)
	cond = sync.NewCond(locker)
)

func test(x int) {
	cond.L.Lock()
	defer cond.L.Unlock()

	cond.Wait()
	zap.S().Info("cond test x -> : ", x)
	time.Sleep(time.Second)
}

func main() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)

	for i:=0; i<40; i++ {
		go test(i)
	}
	zap.S().Info("start all")
	cond.Broadcast()
	time.Sleep(time.Minute)
}