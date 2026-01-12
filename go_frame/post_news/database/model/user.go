package model

import "time"

// User 用户信息表 对应数据库 user 表
// 完美匹配字段+注释+gorm标签+json标签，主键自增/非空/默认值/唯一索引全部适配
type User struct {
	Id         int       `gorm:"primaryKey;autoIncrement;comment:用户id,自增" json:"id"`
	Name       string    `gorm:"type:varchar(20);not null;uniqueIndex:idx_name;comment:用户名" json:"name"`
	Password   string    `gorm:"type:char(32);not null;comment:密码的md5" json:"password"`
	CreateTime time.Time `gorm:"type:datetime;default:current_timestamp;comment:用户注册时间" json:"create_time"`
	UpdateTime time.Time `gorm:"type:datetime;default:current_timestamp;autoUpdateTime;comment:最后修改时间" json:"update_time"`
}

// TableName 指定User结构体对应的数据库表名
func (User) TableName() string {
	return "user"
}
