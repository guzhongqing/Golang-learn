package model

import "time"

// News 新闻信息表 对应数据库 news 表
type News struct {
	Id         int       `gorm:"primaryKey;autoIncrement;comment:新闻id" json:"id"`
	UserId     int       `gorm:"type:int;not null;index:idx_user;comment:发布者id" json:"user_id"`
	Title      string    `gorm:"type:varchar(100);not null;comment:新闻标题" json:"title"`
	Article    string    `gorm:"type:text;not null;comment:正文" json:"article"`
	CreateTime time.Time `gorm:"type:datetime;default:current_timestamp;comment:发布时间" json:"create_time"`
	UpdateTime time.Time `gorm:"type:datetime;default:current_timestamp;autoUpdateTime;comment:最后修改时间" json:"update_time"`
	DeleteTime time.Time `gorm:"type:datetime;default:null;comment:删除时间" json:"delete_time"`
}

// TableName 指定News结构体对应的数据库表名
func (News) TableName() string {
	return "news"
}
