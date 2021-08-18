//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/18 6:02 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		state = make(map[int]int)
		mutex = &sync.Mutex{}
		readOps uint64
		writeOps uint64
	)

	for r:=0; r<100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0; w<10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps : ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps : ", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state : ", state)
	mutex.Unlock()
}