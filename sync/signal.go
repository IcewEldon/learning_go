//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/9/1 2:14 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"learning_go/define"
	"learning_go/log"
	"os"
	"os/signal"
	"syscall"
)

var sigUser = true


func main() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)

	//ignore ctrl-c
	signal.Ignore(os.Interrupt)

	sigChan := make(chan os.Signal, 2)

	//notify SIGHUP
	signal.Notify(sigChan, syscall.SIGHUP)

	//notify SIGUSR1
	signal.Notify(sigChan, syscall.SIGUSR1)

	go func () {
		for {
			switch <-sigChan {
			case syscall.SIGHUP:
				zap.S().Info("sighup, reset signal")
				signal.Reset(syscall.SIGHUP)
			case syscall.SIGUSR1:
				zap.S().Info("signal fall in sigusr1")
				if sigUser {
					zap.S().Info(
						"first usr1, notify interrupt which had ignore")
					usrSigChan := make(chan os.Signal, 1)

					//notify interrupt
					signal.Notify(usrSigChan, os.Interrupt)
					go handlerInterrupt(usrSigChan)
				}
			}
		}
	}()

	select {}
}

func handlerInterrupt(c <- chan os.Signal) {
	for {
		switch <- c {
		case os.Interrupt:
			zap.S().Info("handlerInterrupt : signal interrupt")
		}
	}
}