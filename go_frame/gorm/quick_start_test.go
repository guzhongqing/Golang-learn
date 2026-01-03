package gorm

import (
	"fmt"
	"testing"
)

func TestInitGorm(t *testing.T) {
	config := InitViper("../conf", "mysql", "yaml")
	db := InitGorm(config)
	if db == nil {
		t.Fatal("mysql 连接失败")
	}

	// 插入数据
	login := Login{
		Username: "admin",
		Password: "123456",
	}
	// 这里login表没有主键和唯一索引，所以可以重复插入相同数据，数据表里面会有多行相同数据
	// 但是在查询时，只能查询到第一个插入的数据
	db.Create(&login)
	// 查询数据
	var login2 Login
	db.Find(&login2)
	if login2.Username != "admin" {
		t.Fatal("查询数据失败")
	} else {
		fmt.Printf("查询的数据%v\n", login2)
	}
}
