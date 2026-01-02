package redis

import (
	"context"
	"go_frame/io" // 使用模块相对路径导入io包
	"testing"
)

var config = InitRedisViper("../conf", "redis", io.YAML)
var client = InitRedis(config)

func TestInitRedis(t *testing.T) {
	if client == nil {
		t.Errorf("InitRedis failed")
	}
}

func TestStringValue(t *testing.T) {
	ctx := context.Background()
	stringValue(ctx, client)
}

func TestDeleteKey(t *testing.T) {
	ctx := context.Background()
	DeleteKey(ctx, client)
}

func TestWriteStruct2Redis(t *testing.T) {
	ctx := context.Background()
	WriteStruct2Redis(ctx, client)
}

func TestListValue(t *testing.T) {
	ctx := context.Background()
	listValue(ctx, client)
}

func TestSetValue(t *testing.T) {
	ctx := context.Background()
	setValue(ctx, client)
}

func TestScanKeys(t *testing.T) {
	ctx := context.Background()
	scanKeys(ctx, client)
}
