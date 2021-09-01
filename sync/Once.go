//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 4:23 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"learning_go/define"
	"learning_go/log"
	"sync"
	"time"
)

func main() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)
	var once sync.Once
	onceBody := func () {
		zap.S().Info("something just want do once")
	}

	go func() {
		zap.S().Info("Once in function 1")
		once.Do(onceBody)
	}()

	go func() {
		zap.S().Info("Once in function 2")
		once.Do(onceBody)
	}()

	go func() {
		zap.S().Info("Once in function 3")
		once.Do(onceBody)
	}()
	time.Sleep(3e9)
}
