package pkg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func init() {
	var coreList []zapcore.Core
	//日志格式
	ec := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006/01/02 - 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	hight := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	low := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zap.ErrorLevel && l > zap.DebugLevel
	})
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log",
		MaxSize:    10,   //文件大小，单位MB
		MaxBackups: 5,    //最大保留日志文件数量
		MaxAge:     30,   //日志保留天数
		Compress:   true, //是否压缩
	})
	infoCore := zapcore.NewCore(zapcore.NewConsoleEncoder(ec), zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), low)
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    5,    //文件大小，单位MB
		MaxBackups: 5,    //最大保留日志文件数量
		MaxAge:     60,   //日志保留天数
		Compress:   true, //是否压缩
	})
	errorCore := zapcore.NewCore(zapcore.NewConsoleEncoder(ec), zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), hight)
	coreList = append(coreList, infoCore)
	coreList = append(coreList, errorCore)
	//addcaller() 显示行号和文件名
	Log = zap.New(zapcore.NewTee(coreList...), zap.AddCaller())
}
