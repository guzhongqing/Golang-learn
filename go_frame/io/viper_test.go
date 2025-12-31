package io

import (
	"fmt"
	"testing"
	"time"
)

func TestViper1(t *testing.T) {
	// 读取方式一
	dbViper := InitViper("../conf", "mysql", YAML)
	if !dbViper.IsSet("blog.port") {
		fmt.Println("blog.port 不存在")
	} else {
		port := dbViper.GetInt("blog.port")
		fmt.Printf("port:%d\n", port)
	}

}

func TestViper2(t *testing.T) {
	// 读取方式二
	dbViper := InitViper("../conf", "mysql", YAML)
	dbViper.WatchConfig()
	type MySQLConfig struct {
		Bolg struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Log      string `mapstructure:"log"`
		} `mapstructure:"blog"`
	}

	var mysqlConfig MySQLConfig
	// 这里是把配置文件的内容反序列化到 mysqlConfig 结构体中
	if err := dbViper.Unmarshal(&mysqlConfig); err != nil {
		fmt.Printf("配置文件解析失败: %s \n", err)
		t.Fail()
	} else {
		fmt.Printf("mysqlConfig: %+v \n", mysqlConfig)
	}

	// 验证监听配置文件变化
	for {
		time.Sleep(5 * time.Second)
		if dbViper.IsSet("blog.port") {
			// 不会变化，因为 mysqlConfig 结构体没有更新
			fmt.Printf("mysqlConfig里面的port:%d\n", mysqlConfig.Bolg.Port)
			// 会变化，因为每次都是从配置文件中读取最新的值
			fmt.Printf("viper里面的port:%d\n", dbViper.GetInt("blog.port"))

		}
	}
}
