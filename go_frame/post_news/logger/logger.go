package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 保留原生对象，兼容原有写法（可选使用）
var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

// InitLogger 初始化日志，项目启动时在main.go调用一次即可
// level: 日志级别 debug|info|warn|error ，开发传debug，生产传info
// logFile: 自定义日志文件名 例如"app.log"，传空字符串则默认使用【日期命名】：2026-01-14_app.log
func InitLogger(level string, logFile string) {
	// 1. 获取项目运行根目录（绝对路径，永不跑偏）
	rootDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("获取项目根目录失败: %w", err))
	}

	// 2. 拼接日志文件完整路径，默认按日期命名，传参则自定义
	if logFile == "" {
		logFile = fmt.Sprintf("%s_app.log", time.Now().Format("2006-01-02"))
	}
	logPath := filepath.Join(rootDir, "log", logFile) // 完整文件路径 ./log/xxx.log
	logDir := filepath.Dir(logPath)                   // 提取日志目录 ./log

	// 3. 自动创建日志目录，不存在则创建，权限0755生产标准
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Errorf("创建日志目录失败: %w", err))
	}

	// ========== 日志切割配置（核心，修复BUG1：赋值正确的日志文件路径） ==========
	hook := lumberjack.Logger{
		Filename:   logPath, // ✅ 正确：完整日志文件路径
		MaxSize:    100,     // 单个日志文件最大100MB自动切割
		MaxBackups: 30,      // 保留最多30个旧日志文件
		MaxAge:     7,       // 日志文件保留7天自动清理
		Compress:   true,    // 旧日志自动gzip压缩，节省磁盘
		LocalTime:  true,    // 新增：切割日志的文件名使用本地时间，而非UTC时间
	}

	// ========== zap编码器配置 ==========
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller", // 打印：包名/文件名:行号
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder, // INFO/ERROR 大写
		EncodeTime:     customTimeEncoder,           // 自定义时间格式
		EncodeCaller:   zapcore.ShortCallerEncoder,  // 精简文件路径，不显示绝对路径
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// ========== 设置日志级别 ==========
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// ========== 控制台+文件 双输出 ==========
	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout), // 控制台输出
		zapcore.AddSync(&hook),     // 文件输出
	)

	// ========== 开发/生产 不同编码器 ==========
	var encoder zapcore.Encoder
	if level == "debug" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 开发：彩色易读
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig) // 生产：JSON结构化，日志平台可解析
	}

	// ========== 构建zap日志核心 ==========
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	SugarLogger = Logger.Sugar()

	// 初始化成功日志
	Info("日志初始化成功 ✔️", "日志文件路径", logPath)
}

// customTimeEncoder 自定义日志时间格式 2006-01-02 15:04:05
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// ===================== 核心封装：全局极简调用方法（重点！） =====================
// 所有方法支持 任意参数格式，和 zap.SugarLogger 用法完全一致
// 示例：util.Info("用户注册成功", "userId", 100, "userName", "张三")

// Debug 调试日志 - 开发环境打印，生产环境关闭
func Debug(args ...interface{}) {
	SugarLogger.Debug(args...)
}

// Info 普通日志 - 生产环境核心日志，记录正常业务流程
func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

// Warn 警告日志 - 非错误，需要关注的异常场景（如：用户名已存在、参数为空）
func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

// Error 错误日志 - 程序运行错误，需要排查（如：数据库失败、加密失败、接口调用失败）
func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}

// Panic 致命日志 - 打印日志后直接panic终止程序，仅用于不可恢复的致命错误
func Panic(args ...interface{}) {
	SugarLogger.Panic(args...)
}

// ===================== 带格式化的日志方法（可选，锦上添花） =====================
// 支持 fmt.Sprintf 风格的格式化日志，按需使用
func Infof(format string, args ...interface{}) {
	SugarLogger.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	SugarLogger.Errorf(format, args...)
}

// CloseLogger 程序退出时调用，释放日志资源，刷入缓冲区
func CloseLogger() {
	_ = Logger.Sync()
}
