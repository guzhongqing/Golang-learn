package handler

import (
	"go_frame/post_news/database/gorm"
	"go_frame/post_news/handler/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var reqUser model.LoginUserReq
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := gorm.RegisterUser(reqUser.Name, reqUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	var reqUser model.LoginUserReq
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用登录函数，获取数据库用户模型
	dbUser, err := gorm.LoginUser(reqUser.Name, reqUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// 登录成功，返回cookie
	// 设置cookie
	c.SetCookie("uid", strconv.Itoa(dbUser.ID), 3600, "/", "localhost", false, true)

	// 可以直接返回数据库用户模型，或者转换为响应模型
	c.JSON(http.StatusOK, gin.H{"user": dbUser})
}

// logoutUser 用户退出
func LogOutUser(c *gin.Context) {
	// 清除cookie
	c.SetCookie("uid", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "用户注销成功"})
}

// ModifyPassword 修改密码
func ModifyPassword(c *gin.Context) {
	var reqUser model.ModifyPassReq
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从cookie中获取用户ID
	uidStr, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	// 转换为整数
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户ID格式错误"})
		return
	}
	// 调用修改密码函数
	err = gorm.UpdatePassword(uid, reqUser.NewPassword, reqUser.OldPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})
}
