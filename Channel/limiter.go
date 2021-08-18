//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/18 4:00 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"time"
)

func BurstsLimit() {
	bursts := make(chan time.Time, 3)
	for i:=0; i<3; i++ {
		bursts <- time.Now()
	}
	close(bursts)

	for burst := range bursts{
		fmt.Println("burst : ", burst)
	}
}

func main() {
	requests := make(chan int, 5)
	for i:=1; i<=5;i++ {
		requests <- 1
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<- limiter
		fmt.Println("requests : ", req, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)

	for i:=0; i<3; i++ {
		burstLimiter <- time.Now() // 所谓爆发限制，是提前打满req缓冲
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i:=1; i<=5; i++ {
		burstRequests <- i
	}
	close(burstRequests)
	for req := range burstRequests {
		<- burstRequests
		fmt.Println("request", req, time.Now())
	}
}