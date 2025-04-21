package main

import (
	"os"

	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getLogWriter()
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
