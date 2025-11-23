package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel    = "LOG_LEVEL"
	envLogOutput   = "LOG_OUTPUT"
	envServiceName = "SERVICE_NAME"
)

var (
	log logger
)

type monitoringLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	var err error
	logConfig := zap.NewProductionConfig()
	logConfig.OutputPaths = []string{getOutput()}
	logConfig.Level = zap.NewAtomicLevelAt(getLevel())
	logConfig.Encoding = "json"

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "timestamp"
	encodeConfig.LevelKey = "level"
	encodeConfig.MessageKey = "msg"
	encodeConfig.StacktraceKey = ""
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logConfig.EncoderConfig = encodeConfig

	if log.log, err = logConfig.Build(
		zap.AddCallerSkip(1),
		zap.Fields(zap.String("service", getServiceName())),
	); err != nil {
		panic(err)
	}
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	output := strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}

func getServiceName() string {
	serviceName := strings.TrimSpace(os.Getenv(envServiceName))
	if serviceName == "" {
		return "unknown"
	}
	return serviceName
}

func GetLogger() monitoringLogger {
	return log
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
