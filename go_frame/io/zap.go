package io

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// lumberjackv2 "gopkg.in/natefinch/lumberjack.v2"
)

func SayHello() {
	fmt.Println("Hello, World!")
}

func InitZap1(logFile string) *zap.Logger {
	fmt.Println("InitZap1 called")
	// logger := zap.NewExample() // 测试环境日志
	// logger, err := zap.NewDevelopment() // 开发环境日志
	// logger, _ := zap.NewProduction() // 生产环境日志
	//
	encoderConfig := zap.NewProductionEncoderConfig()
	// 指定时间展示格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// 修改默认ts字段名为time
	encoderConfig.TimeKey = "time"
	// 日志级别全字母大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	// if err != nil {
	// 	fmt.Println("Open log file failed, err:", err)
	// 	return nil
	// }

	// 按照文件大小进行日志轮转
	// lumberjackLogger := &lumberjackv2.Logger{
	// 	Filename:   logFile,
	// 	MaxSize:    10,   // 每个日志文件最大10MB
	// 	MaxBackups: 5,    // 最多保留5个备份
	// 	MaxAge:     30,   // 最多保留30天的日志文件
	// 	Compress:   true, // 是否压缩旧日志文件
	// }

	// 按照时间进行日志轮转
	rotateOut, err := rotatelogs.New(
		logFile+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logFile),         // 软链接，指向最新日志文件
		rotatelogs.WithRotationTime(1*time.Hour), // 每小时轮转一次
		rotatelogs.WithMaxAge(24*time.Hour),      // 最多保留24小时的日志文件
	)
	if err != nil {
		fmt.Println("Create rotatelogs failed, err:", err)
		return nil
	}

	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(encoderConfig), // 设置json格式输出
		zapcore.NewConsoleEncoder(encoderConfig), // 设置控制台格式输出
		// zapcore.AddSync(file),                    // 写入文件
		// zapcore.AddSync(lumberjackLogger), // 使用lumberjack进行日志轮转
		zapcore.AddSync(rotateOut), // 使用rotatelogs进行日志轮转
		zapcore.DebugLevel,
	)
	logger := zap.New(
		core,
		zap.AddCaller(),                       // 打印调用者, 包含文件名、行号
		zap.AddStacktrace(zapcore.ErrorLevel), // 调用栈跟踪, Error级别以上
		zap.Hooks(func(entry zapcore.Entry) error {
			fmt.Println("AddHook called, entry:", entry)
			if entry.Level >= zapcore.ErrorLevel {
				fmt.Printf("Error log: %s\n", entry.Message)
			}
			return nil
		}), // 添加沟子
	)
	// 添加公共字段
	logger = logger.With(
		zap.Namespace("公共字段"),
		zap.String("app", "myApp"),
		zap.String("env", "production"),
	)

	return logger

}

// 简化配置版的InitZap1，少了日志轮转，调用栈，钩子函数等配置
func InitZap2(logFile string, logLevel zapcore.Level) *zap.Logger {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Encoding:         "json",                                  // 输出格式（json/console）
		OutputPaths:      []string{"stdout", logFile},             // 正常日志输出路径（终端+文件）
		ErrorOutputPaths: []string{"stderr"},                      // 错误日志终端输出
		InitialFields:    map[string]interface{}{"biz": "search"}, // 全局公共字段
		EncoderConfig: zapcore.EncoderConfig{
			EncodeTime:  zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
			TimeKey:     "time",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}
	logger, _ := config.Build()
	defer logger.Sync()
	logger = logger.With(
		zap.Namespace("公共字段"),
		zap.String("app", "myApp"),
		zap.String("env", "production"),
	)
	return logger

}
