package gorm

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// gorm默认情况，如果有字段为整型的ID/Id 会默认设置为 自增主键，结构体名的蛇形复数作为表名（所以这里的User会默认表名为users），字段名的蛇形直接作为列名
// 不建议使用AutoMigrate自动迁移
// go的int类型默认对应数据库的int(11)有符合类型,可以用type指定int unsigned
// go的string类型默认对应数据库的varchar(255)类型,可以用type指定char,text,longtext等
// go的time.Time类型默认对应数据库的datetime类型,可以用type指定date,timestamp

type User struct {
	Id     int // gorm建表时会自动给ID/Id设置「自增主键」，不需要手动设置
	UserId int `gorm:"column:uid"` // 对应数据库的uid字段
	Degree string
	// GORM 对结构体字段名的蛇形命名（snake_case） 转换规则是：大驼峰拆分为小写 + 下划线分隔，连续大写缩写（如 JSON）整体转小写
	// 所以默认 KeywordsJSONArray → 蛇形字段名：keywords_json_array
	// 所以默认KeywordsJSONObject → 蛇形字段名：keywords_json_object
	KeywordsJSONArray  datatypes.JSONSlice[string] // 对应数据库的json数组类型，默认为text类型
	KeywordsJSONObject datatypes.JSONMap           // 对应数据库的json对象类型，默认为text类型

	// gorm默认 CreatedAt和UpdatedAt 字段会自动设置为当前创建和更新时间
	// 如果数据库设置了CURRENT_TIMESTAMP,也是gorm填充当前时间写入数据库
	CreatedAt time.Time `gorm:"column:create_time"`
	// time.Time 对应数据库的 datetime 类型（精确到时分秒），type 可以指定数据库的类型，date对应数据库的日期类型（精确到年月日）
	// CreatedAt time.Time `gorm:"column:create_time; type:date"`
	// 如果数据库设置了CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,也是gorm填充当前时间写入数据库
	UpdatedAt time.Time `gorm:"column:update_time"`
	Gender    string
	City      string
	Province  string `gorm:"-"` // 结构体有，但是数据库没有，所以需要忽略
}

// TableName 实现 gorm.Tabler 接口，指定数据库表名
func (User) TableName() string {
	return "user"
}

// 注意：数据库varcahr(n)和char(n)类型，其中n为字符数（如，'男'为1个字符），长度大于n会报错
// char(n)存入时自动补空格到 n 长度，读取时自动剔除所有尾部空格。
// varchar(n) 存入 / 读取时保留原始字符串尾部空格

func Create(db *gorm.DB) {
	// 插入数据
	user := User{
		UserId: rand.Intn(100000),
		Degree: "Bachelor",

		// 原生类型可以直接赋值给「底层类型相同」的自定义类型（隐式转换），所以这里两种方式都可以
		// KeywordsJSONArray: []string{"key1", "key2"},
		KeywordsJSONArray: datatypes.JSONSlice[string]{"key1", "key2"},

		// KeywordsJSONObject: map[string]any{"key1": "value1", "key2": "value2"},
		KeywordsJSONObject: datatypes.JSONMap{"key1": "value1", "key2": "value2"},

		Gender: "Male",
		City:   "Shanghai",
	}

	// 插入数据
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("插入数据失败:", result.Error)
	} else {
		fmt.Printf("插入数据成功: %#v\n", user)
	}

	// 查询所有数据
	var users []User
	result = db.Find(&users)
	if result.Error != nil {
		fmt.Println("查询数据失败:", result.Error)
	} else {
		fmt.Printf("查询数据成功: %#v\n", users)
	}

}
