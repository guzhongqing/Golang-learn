package io_test

import (
	"go_frame/io"
	"testing"

	"go.uber.org/zap"
)

func TestSayHello(t *testing.T) {
	io.SayHello()
}

func TestInitZap1(t *testing.T) {
	logger := io.InitZap1("app.log")
	defer logger.Sync() // 同步内容到文件中
	logger.Debug("debug")
	logger.Info("This is a test log message from TestInitZap1", zap.Int("ageKey", 42))
	logger.Warn("warn", zap.Namespace("nameSpaceKey"), zap.String("nameKey", "nameValue"))
	logger.Error("error message")

	// 格式化输出性能降低50%左右
	zapSugar := logger.Sugar()
	zapSugar.Infof("infof: %f", 3.14)
}

func TestInitZap2(t *testing.T) {
	logger := io.InitZap2("app.log", zap.DebugLevel)
	defer logger.Sync() // 同步内容到文件中
	logger.Debug("debug")
	logger.Info("This is a test log message from TestInitZap1", zap.Int("ageKey", 42))
	logger.Warn("warn", zap.Namespace("nameSpaceKey"), zap.String("nameKey", "nameValue"))
	logger.Error("error message")

	// 格式化输出性能降低50%左右
	zapSugar := logger.Sugar()
	zapSugar.Infof("infof: %f", 3.14)
}
