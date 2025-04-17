package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func configLogger(config settings.LoggerSetting) lumberjack.Logger {
	return lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   // days
		Compress:   config.Compress, // disabled by default
	}
}

func NewLogger(config settings.LoggerSetting) *LoggerZap {
	encoder := getEncoderLog()
	logLevel := config.LogLevel
	level := getLogLevel(logLevel)
	hook := configLogger(config)
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{
		zap.New(
			core,
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		),
	}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"

	return zapcore.NewJSONEncoder(encodeConfig)
}
