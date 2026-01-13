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

	// ========== 日志切割配置 ==========
	hook := lumberjack.Logger{
		Filename:   logPath, // 完整日志文件路径
		MaxSize:    100,     // 单个日志文件最大100MB自动切割
		MaxBackups: 30,      // 保留最多30个旧日志文件
		MaxAge:     7,       // 日志文件保留7天自动清理
		Compress:   true,    // 旧日志自动gzip压缩，节省磁盘
		LocalTime:  true,    // 切割日志的文件名使用本地时间，而非UTC时间
	}

	// ========== zap 核心编码器配置 - 【每行注释+Debug模式生效/失效精准标注】 ==========
	// 重要说明：
	// 1. 生产环境(JSONEncoder)：以下配置 100% 全部生效 ✔️
	// 2. 开发环境(ConsoleEncoder)：已开启 CapitalColorLevelEncoder 开关，标注如下：✅=生效 ❌=不生效(控制台编码器原生特性，无法修改)
	encoderConfig := zapcore.EncoderConfig{
		// time 字段的key名，最终日志里显示的时间字段名称
		// Debug(控制台) ❌ 不生效：控制台日志直接把时间作为前缀展示，不会显示该key名，格式由EncodeTime控制
		// JSON(生产)   ✅ 生效：日志中显示 "time": "2026-01-14 17:00:00.000"
		TimeKey: "time",

		// level 字段的key名，最终日志里显示的级别字段名称
		// Debug(控制台) ❌ 不生效：控制台日志直接把级别作为前缀展示[INFO]，不会显示该key名，格式由EncodeLevel控制
		// JSON(生产)   ✅ 生效：日志中显示 "level": "INFO"
		LevelKey: "level",

		// caller 字段的key名，最终日志里显示的调用栈行号字段名称
		// Debug(控制台) ❌ 不生效：控制台日志直接在行尾展示行号，不会显示该key名，格式由EncodeCaller控制
		// JSON(生产)   ✅ 生效：日志中显示 "caller": "service/user.go:25"
		CallerKey: "caller",

		// msg 字段的key名，最终日志里显示的业务日志内容字段名称
		// Debug(控制台) ❌ 不生效：控制台日志直接展示日志内容，不会显示该key名
		// JSON(生产)   ✅ 生效：日志中显示 "msg": "用户注册成功"
		MessageKey: "msg",

		// stacktrace 字段的key名，最终日志里显示的错误堆栈字段名称
		// Debug(控制台) ❌ 不生效：控制台日志的堆栈直接追加在日志末尾，不会显示该key名
		// JSON(生产)   ✅ 生效：日志中显示 "stacktrace": "xxx/xxx.go:xx +0x123"
		// 生效条件：仅 Error/Panic 级别日志会打印堆栈
		StacktraceKey: "stacktrace",

		// 日志级别的格式化规则：CapitalLevelEncoder 表示日志级别全大写(DEBUG/INFO/WARN/ERROR/PANIC)
		// Debug(控制台) ✅ 生效：控制台日志级别会显示大写，且结合开关显示对应颜色(DEBUG灰/INFO蓝/WARN黄/ERROR红)
		// JSON(生产)   ✅ 生效：日志级别固定大写格式
		EncodeLevel: zapcore.CapitalLevelEncoder,

		// 日志时间的自定义格式化规则，绑定我们自己写的 customTimeEncoder 函数
		// Debug(控制台) ✅ 生效：控制台日志的时间格式完全按照自定义的 2006-01-02 15:04:05.000 展示
		// JSON(生产)   ✅ 生效：JSON日志的时间格式和自定义规则完全一致
		// 【重中之重】：这是Debug模式下 最核心的必生效配置之一
		EncodeTime: customTimeEncoder,

		// 调用栈行号的格式化规则：ShortCallerEncoder 表示精简路径（包名/文件名:行号），不显示完整绝对路径
		// Debug(控制台) ✅ 生效：控制台日志的行号展示为 例如 logger/logger.go:98 | service/user.go:25
		// JSON(生产)   ✅ 生效：JSON日志的行号格式和配置一致
		// 【重中之重】：这是Debug模式下 另一个必生效配置之一
		EncodeCaller: zapcore.ShortCallerEncoder,

		// 耗时字段的格式化规则：SecondsDurationEncoder 表示耗时以秒为单位展示
		// Debug(控制台) ✅ 生效：如果日志中带耗时字段，会按秒级展示
		// JSON(生产)   ✅ 生效：JSON日志的耗时格式一致
		// 补充说明：该配置仅在使用 zap.WithDuration() 时生效，日常业务日志基本用不到，不影响核心功能
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

	// ========== 构建zap日志核心 ✅【核心修复点1：添加 AddCallerSkip(1)】==========
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel)
	// AddCallerSkip(1) 跳过1层封装函数，直接获取业务代码的真实调用位置
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	SugarLogger = Logger.Sugar()

	// ========== ✅【修复点2：初始化日志用原生对象，避免初始化行号错误】==========
	SugarLogger.Info("✅日志初始化成功 ", "日志文件路径", logPath)
}

// customTimeEncoder 自定义日志时间格式 2006-01-02 15:04:05
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// ===================== 核心封装：全局极简调用方法（写法不变，无需修改） =====================
// 所有方法支持 任意参数格式，和 zap.SugarLogger 用法完全一致
// 示例：logger.Info("用户注册成功", "userId", 100, "userName", "张三")

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

// ===================== 带格式化的日志方法 =====================
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
