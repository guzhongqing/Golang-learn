package gorm

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	config := InitViper("../conf", "mysql", "yaml")
	db := CreateConnection(config)
	if db == nil {
		t.Fatal("mysql 连接失败")
	}
	Create(db)

}
