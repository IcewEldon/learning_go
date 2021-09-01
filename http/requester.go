//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/27 2:36 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"bufio"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"learning_go/log"
	"net/http"
)

func main() {
	log.InitLogger("", zapcore.InfoLevel.String())
	responser, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer responser.Body.Close()

	zap.S().Info("response status : ", responser.Status)

	scanner := bufio.NewScanner(responser.Body)
	for i:=0; scanner.Scan() && i<5; i++ {
		zap.S().Info("Do scanner : ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
