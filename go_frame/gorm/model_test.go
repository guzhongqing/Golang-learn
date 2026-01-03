package gorm

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var config *viper.Viper
var db *gorm.DB

func Init() {
	config := InitViper("../conf", "mysql", "yaml")
	db = InitGorm(config)
	if db == nil {
		fmt.Println("mysql 连接失败")
	}

}

func TestCreate(t *testing.T) {
	Init()
	Create(db)
}
