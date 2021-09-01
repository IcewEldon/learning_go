//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/27 2:38 下午
//Copyright 2021 www.ucloud.cn

package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func InitLogger(path, level string){

	defaultLevtl := zapcore.InfoLevel
	switch strings.ToLower(level) {
	case "debug":
		defaultLevtl = zapcore.DebugLevel
	case "warn":
		defaultLevtl = zapcore.WarnLevel
	case "error":
		defaultLevtl = zapcore.ErrorLevel
	case "fatal":
		defaultLevtl = zapcore.FatalLevel
	default:
	}

	// output log setting
	sync := zapcore.AddSync(os.Stdout)
	if len(path) > 0 {
		lj := &lumberjack.Logger{
			Filename: path,
			MaxSize: 10,
			MaxAge: 90,
		}
		sync = zapcore.AddSync(lj)
	}

	core := zapcore.NewCore(GetEncoder(), sync, defaultLevtl)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)

	zap.S().Infof("logger init done")
}

func GetEncoder() zapcore.Encoder{
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}