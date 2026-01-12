package gorm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 全局导出DB变量，供外部所有包直接调用
var DB *gorm.DB

// 初始化GORM日志文件：接收【日志文件名】，自动拼接项目根目录生成绝对路径
func initGormLogFile(logFileName string) (*os.File, error) {
	// 1. 获取项目根目录绝对路径（运行时的工作目录，永远正确）
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取项目根目录失败: %w", err)
	}

	// 2. 核心：拼接路径 = 项目根目录/log/日志文件名
	// 日志会统一放在 项目根目录/log/ 下，整洁规范
	logPath := filepath.Join(rootDir, "log", logFileName)
	logDir := filepath.Dir(logPath)

	// 3. 自动创建日志目录，不存在则创建，存在无操作
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 4. 打开/创建日志文件，标准日志文件权限0644
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件失败: %w", err)
	}

	return logFile, nil
}

// 创建数据库连接 - 最终完整版，无BUG、无冗余、读取配置文件logfile
func CreateConnection(config *viper.Viper) error {
	// 读取mysql节点下的所有数据库配置
	host := config.GetString("mysql.host")
	port := config.GetInt("mysql.port")
	username := config.GetString("mysql.username")
	password := config.GetString("mysql.password")
	dbname := config.GetString("mysql.dbname")

	// ✅ 修改点1：读取 mysql.logfile 配置项 【核心】
	logFileName := config.GetString("mysql.logfile")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// ✅ 修改点2：把读取到的日志文件名传入日志初始化函数
	logFile, err := initGormLogFile(logFileName)
	if err != nil {
		log.Fatalf("初始化gorm日志文件失败: %v", err)
		return err
	}

	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:        500 * time.Millisecond,
			LogLevel:             logger.Info,
			ParameterizedQueries: false,
			Colorful:             false, // 写入文件必须关闭彩色，防止乱码
		},
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 匹配你的user/news单数表名，不用改
		},
		Logger:               newLogger,
		DryRun:               false, // 必改的BUG，已修复，正常执行SQL
		DisableAutomaticPing: false,
	})
	if err != nil {
		log.Fatalf("mysql 连接失败: %v", err)
		return err
	}

	// 赋值给全局DB变量，供外部调用
	DB = conn

	log.Printf("✅ mysql 连接成功: %s", dsn)
	// 设置数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败: %v", err)
		return err
	}
	log.Printf("✅ 数据库连接池配置: 最大打开连接数=100, 最大空闲连接数=10, 连接最大生命周期=1小时")
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 数据库健康检查
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("数据库健康检查失败: %v", err)
		return err
	}
	log.Printf("✅ 数据库健康检查通过")

	// 无返回值，成功返回nil
	return nil
}
