package util

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

const (
	JSON = "json"
	YAML = "yaml"
	ENV  = "env"
)

func InitViper(dir, file, FileType string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(dir)
	config.SetConfigName(file)
	config.SetConfigType(FileType)

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件 %s 解析失败: %s \n", path.Join(dir, file+"."+FileType), err))
	}

	return config
}
