package gorm_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go_frame/post_news/database/gorm"
	"go_frame/post_news/logger"
	"go_frame/post_news/util"
	"testing"
)

func init() {
	fmt.Println("Hello, World!")
	// 初始化配置

	// 初始化日志
	logger.InitLogger("debug", "logger.test.log")
	defer logger.CloseLogger() // 程序退出释放资源

	// 初始化数据库
	viperConfig := util.InitViper("../../conf", "mysql", util.YAML)
	// 初始化数据库连接,相关路径是相对于main.go的路径
	// 2. 初始化数据库连接（带错误捕获）
	if err := gorm.CreateConnection(viperConfig); err != nil {
		logger.Error("数据库初始化失败", "err", err)
	}
}
func hashPassword(password string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(password))
	digest := md5Hash.Sum(nil)
	// md5 输出128bit，16个byte元素，按照16进制（4位）表示，共32个字符，4bit表示一个字符
	return hex.EncodeToString(digest)

}

func TestRegisterUser(t *testing.T) {
	// 测试注册用户
	// name := "testuser"
	name := "agufish"
	// name := "user"
	// password := "123456"
	password := "654321"

	// 模拟密码加密
	hashedPassword := hashPassword(password)
	id, err := gorm.RegisterUser(name, hashedPassword)
	if err != nil {
		logger.Errorf("注册用户失败，ID：%d\n", id)
		t.Fatalf("注册用户失败：%v", err)

	}
	if id <= 0 {
		logger.Error("注册用户返回无效ID：%d", id)
		t.Fatalf("注册用户返回无效ID：%d", id)
	}
	// 打印注册成功日志
	logger.Infof("注册用户成功，ID：%d\n", id)
}

func TestLogOffUser(t *testing.T) {
	// 测试注销用户
	id := 2
	err := gorm.LogOffUser(id)
	if err != nil {
		logger.Errorf("注销用户失败，ID：%d\n", id)
		t.Fatalf("注销用户失败：%v", err)
	}
	// 打印注销成功日志
	logger.Infof("注销用户成功，ID：%d\n", id)
}

func TestUpdatePassword(t *testing.T) {
	// 测试更新密码
	id := 2
	oldPassword := "123456"
	newPassword := "654321"
	err := gorm.UpdatePassword(id, newPassword, hashPassword(oldPassword))
	if err != nil {
		logger.Errorf("更新密码失败，ID：%d\n", id)
		t.Fatalf("更新密码失败：%v", err)
	}
	// 打印更新成功日志
	logger.Infof("更新密码成功，ID：%d\n", id)
}
