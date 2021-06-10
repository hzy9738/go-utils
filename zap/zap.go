package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

func init() {
	//日志文件名称
	fileName := "goproject.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			Filename: fileName, //文件名称
			MaxSize:  521,      //MB
			//MaxAge:     0,
			MaxBackups: 0,    //最大备份
			LocalTime:  true, //
			Compress:   true, //是否启动压缩
		},
	)
	//编码
	encoder := zap.NewProductionEncoderConfig()
	//时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		zap.NewAtomicLevelAt(zap.DebugLevel))
	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args)
}
