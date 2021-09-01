//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/30 11:49 AM
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool)
	//signal.Notify :
	signal.Notify(signals, syscall.SIGUSR1)

	go func() {
		sig := <-signals
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<- done
	fmt.Println("exiting")
}
