package io

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

// FileType 配置文件类型，如：yaml, json, toml 等
const (
	YAML = "yaml"
	JSON = "json"
	ENV  = "env"
)

func InitViper(dir, file, FileType string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(dir)
	config.SetConfigName(file)     // 文件名，不带路径，不带后缀
	config.SetConfigType(FileType) // 文件后缀类型

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件 %s 解析失败: %s \n", path.Join(dir, file+"."+FileType), err))
	}

	return config
}
