package io

import (
	"fmt"
	"log/slog"
	"os"
)

var logFile = "data/slog.log"

// 定制一个Slogger，使用log/slog包
func NewSlogger() *slog.Logger {
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
	}
	// 使用JSON格式
	// logger := slog.New(slog.NewJSONHandler(fout, &slog.HandlerOptions{
	// 	AddSource: true,
	// 	Level:     slog.LevelInfo,
	// }))

	// 使用文本格式
	logger := slog.New(slog.NewTextHandler(fout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
	return logger

}

func Slog(logger *slog.Logger) {
	// 写入文本
	logger.Debug("加法运算", "a", 1, "b", 2, "sum:", 3)
	logger.Info("加法运算", "a", 1, "b", 2, "sum:", 3)
	logger.Warn("加法运算", "a", 1, "b", 2, "sum:", 3)
	logger.Error("加法运算", "a", 1, "b", 2, "sum:", 3)

	// 控制台输出
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("加法运算", "a", 1, "b", 2, "sum:", 3)
	slog.Info("加法运算", "a", 1, "b", 2, "sum:", 3)
	slog.Warn("加法运算", "a", 1, "b", 2, "sum:", 3)
	slog.Error("加法运算", "a", 1, "b", 2, "sum:", 3)

}
