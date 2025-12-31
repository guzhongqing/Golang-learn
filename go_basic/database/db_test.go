package database_test

import (
	"database/sql"
	"fmt"
	"golang_learn/go_basic/database"
	"testing"

	_ "github.com/go-sql-driver/mysql" //初始化会调用里面的init()函数
)

var db *sql.DB

func init() {
	// 这里的db是全局变量，err是局部变量
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestSelect(t *testing.T) {
	database.Query(db)
}
