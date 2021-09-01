//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/27 4:04 下午
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	zap.S().Info(w, "hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			zap.S().Infof("%v : %v : %v", w, name, h)
		}
	}
}

func main(){
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	zap.S().Info("starting http server ...")
	http.ListenAndServe(":8090", nil)
}