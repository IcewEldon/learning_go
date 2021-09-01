//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 2:28 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"go.uber.org/zap"
	"learning_go/define"
	"learning_go/log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	log.InitLogger(define.DEFAULTPATH, define.INFO)
	if len(os.Args) < 2 {
		zap.S().Infof("Usage: %s [command]", os.Args[0])
		os.Exit(1)
	}

	cmdName := os.Args[1]
	if filepath.Base(os.Args[1]) == os.Args[1] {
		if lp, err := exec.LookPath(os.Args[1]); err != nil {
			zap.S().Infof("look for path error : ", err)
			os.Exit(1)
		} else {
			cmdName = lp
			zap.S().Info("'exec.LookPath find path : ", cmdName)
		}
	}

	procAttr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}

	cwd, err := os.Getwd()
	if err != nil {
		zap.S().Infof("look path error : ", err)
		os.Exit(1)
	}

	start := time.Now()
	process, err := os.StartProcess(cmdName, []string{cwd}, procAttr)
	if err != nil {
		zap.S().Info("start process err : ", err)
		os.Exit(2)
	}

	processState, err := process.Wait()
	if err != nil {
		zap.S().Info("wait error : ", err)
		os.Exit(3)
	}

	zap.S().Info("real -> ", time.Now().Sub(start))
	zap.S().Info("user -> ", processState.UserTime())
	zap.S().Info("system -> ", processState.SystemTime())
}
