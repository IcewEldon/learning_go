//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/27 5:31 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"io/ioutil"
	"learning_go/define"
	"learning_go/log"
	"os/exec"
)

func main() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	zap.S().Info("> date ")
	zap.S().Info(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\n"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	zap.S().Info("> grep hello")
	zap.S().Info(string(grepBytes))

}

