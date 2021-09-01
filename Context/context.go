//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/27 5:12 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"learning_go/log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	zap.S().Info("server : htllo handler started")

	select {
	case <- time.After(10 * time.Second) :
		zap.S().Info(w, "Hello")
	case <- ctx.Done():
		err := ctx.Err()
		zap.S().Info("server : ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main () {
	log.InitLogger("", zapcore.InfoLevel.String())
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
