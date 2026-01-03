package gorm

import (
	"fmt"
	"log/slog"
	"path"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TableName 实现 gorm.Tabler 接口，指定数据库表名
func (Login) TableName() string {
	return "login"
}

func InitViper(dir, file, FileType string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(dir)
	config.SetConfigName(file)
	config.SetConfigType(FileType)

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件 %s 解析失败: %s \n", path.Join(dir, file+"."+FileType), err))
	}

	return config
}

func InitGorm(config *viper.Viper) *gorm.DB {
	host := config.GetString("mysql.host")
	port := config.GetInt("mysql.port")
	username := config.GetString("mysql.username")
	password := config.GetString("mysql.password")
	dbname := config.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		slog.Error("mysql 连接失败", "err", err)
		return nil
	} else {
		slog.Info("mysql 连接成功", "dsn", dsn)
		return db
	}
}
