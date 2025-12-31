package io

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyDate time.Time

// 实现 stringer 接口
func (d MyDate) String() string {
	return time.Time(d).Format(MyDateFormat)
}

const MyDateFormat = "2006-01-02 15:04:05"

func (d MyDate) MarshalJSON() ([]byte, error) {
	// 格式化日期为字符串，把d强转为time.Time类型
	s := fmt.Sprintf("\"%s\"", time.Time(d).Format(MyDateFormat))
	// s 的内容为 "2006-01-02 15:04:05"
	return []byte(s), nil
}

func (d *MyDate) UnmarshalJSON(s []byte) error {
	t, err := time.ParseInLocation(`"`+MyDateFormat+`"`, string(s), time.Local)
	if err != nil {
		return err
	}
	*d = MyDate(t)
	return nil
}

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	heigth   float32
	Birthday MyDate
	Sex      int `json:"gender"`
}

func JsonSerialize() {

	// 实例化一个User结构体
	user := User{
		Name:     "张三",
		Age:      30,
		heigth:   1.8,
		Birthday: MyDate(time.Now()),
		Sex:      1,
	}

	// 序列化结构体为JSON字符串
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("序列化结构体失败: %v\n", err)
		return
	}
	fmt.Printf("序列化后的JSON字符串: %s\n", jsonBytes)

	var u User
	// 反序列化JSON字符串为结构体
	err = json.Unmarshal(jsonBytes, &u)
	if err != nil {
		fmt.Printf("反序列化JSON字符串失败: %v\n", err)
		return
	}
	fmt.Printf("反序列化后的结构体+v: %+v\n", u)
	fmt.Printf("反序列化后的结构体#v: %#v\n", u)

}
