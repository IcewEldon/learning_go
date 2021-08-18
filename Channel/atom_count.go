//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/18 5:51 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		ops uint64
		wg sync.WaitGroup
	)
	for i:=0; i<50; i++ {
		wg.Add(1)

		go func() {
			for c:=0; c<1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops : ", ops)
}
