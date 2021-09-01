//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 4:31 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"learning_go/define"
	"learning_go/log"
	"sync"
	"time"
)

func WaitGroup() {
	wp := new(sync.WaitGroup)
	wp.Add(10)

	for i:=0; i<10; i++ {
		zap.S().Infof("done -> : %d", i)
		wp.Done()
	}

	wp.Wait()
	zap.S().Info("wait done")
}

func Cond() {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	done := false

	cond.L.Lock()
	defer cond.L.Unlock()

	go func() {
		time.Sleep(2e9)
		done = true
		cond.Signal()
	}()

	if !done {
		cond.Wait()
	}

	zap.S().Info("now done is -> : ", done)
}

func main() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)
	WaitGroup()
	Cond()
}