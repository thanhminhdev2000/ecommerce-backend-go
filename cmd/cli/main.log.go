package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "TipsGo", 40)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 40))

	// 2.
	// logger := zap.NewExample()
	// logger.Info(("Hello NewExample"))


	// logger, _ = zap.NewDevelopment()
	// logger.Info(("Hello NewDevelopment"))

	// logger, _ = zap.NewProduction()
	// logger.Info(("Hello NewProduction"))

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1736417862.5964086 -> 2025-01-09T17:17:42.595+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	// ts -> Time
	encodeConfig.TimeKey = "time"
	
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	
	// "caller":"cli/main.log.go:22"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	// os.O_RDWR : read and write
	// os.O_RDONLY: read only
	// os.O_WRONLY : write only
	// os.ModePerm : 0777 max permission
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY , os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}