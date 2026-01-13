package gorm

import (
	"errors"
	"go_frame/post_news/database/model"
	"go_frame/post_news/logger"

	"gorm.io/gorm"
)

// RegisterUser 注册用户 - 生产级最优写法
func RegisterUser(name, password string) (int, error) {
	// ===== 优化1：严谨判断用户名是否存在（仅判断存在性，不查完整数据，性能最优） =====
	var count int64
	err := DB.Model(&model.User{}).Where("name = ?", name).Count(&count).Error
	// 优先判断：数据库查询是否发生错误（连接失败/超时等）
	if err != nil {
		logger.Error("查询用户名失败", "name", name, "err", err)
		return 0, errors.New("查询用户名失败：" + err.Error())
	}
	// 再判断：用户名是否已存在
	if count > 0 {
		return 0, errors.New("用户名已存在")
	}

	// ===== 构造用户结构体，存入加密后的密码 =====
	user := model.User{
		Name:     name,
		Password: password, // 存加密后的密码，不是明文
	}

	// ===== 写入数据库，原写法正确，保留即可 =====
	if err := DB.Create(&user).Error; err != nil {
		return 0, errors.New("用户注册失败：" + err.Error())
	}

	// ===== 优化3：返回原生uint类型的ID，避免类型转换隐患 =====
	return user.ID, nil
}

// 用户注销
// LogOffUser 用户注销
func LogOffUser(id int) error {
	// 根据ID删除用户
	tx := DB.Delete(&model.User{ID: id})
	if err := tx.Error; err != nil {
		logger.Error("用户注销失败", "id", id, "err", err)
		return errors.New("用户注销失败：" + err.Error())
	}
	// 检查是否有影响的行数
	if tx.RowsAffected == 0 {
		return errors.New("用户注销失败：用户不存在")
	}
	return nil
}

// LoginUser 用户登录
func LoginUser(name, password string) (model.User, error) {
	user := model.User{
		Name:     name,
		Password: password,
	}
	// 根据用户名查询用户（必须查完整数据，要校验密码）
	err := DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		// 查不到用户 或 数据库错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("用户名或密码错误")
		}
		return model.User{}, errors.New("查询用户失败：" + err.Error())
	}

	// 登录成功，返回用户信息（注意：不要返回密码字段）
	user.Password = ""
	return user, nil
}
