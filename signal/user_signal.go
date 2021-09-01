//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/30 12:29 PM
//Copyright 2021 www.ucloud.cn

package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//用json配置测试
type Config struct {
	Test1 string `json:"Test1:`
	Test2 int  `json:"Test1:`
}

var (
	config   *Config
	configLock = new(sync.RWMutex)
)

func loadConfig() bool {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("load config error: ", err)
		return false
	}

	//不同的配置规则，解析复杂度不同
	temp := new(Config)
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println("Para config failed: ", err)
		return false
	}

	configLock.Lock()
	config = temp
	configLock.Unlock()
	return true
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func init() {
	if !loadConfig() {
		os.Exit(1)
	}

	//热更新配置可能有多种触发方式，这里使用系统信号量sigusr1实现
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			log.Println("Reloaded config:", loadConfig())
		}
	}()
}

func main() {
	fmt.Println("wait for usr signal")
	select {}
}