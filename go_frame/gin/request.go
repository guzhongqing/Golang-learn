package main

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 查询参数
func QueryGet(c *gin.Context) {
	// 查询参数
	// 127.0.0.1:8080/student?name=张三
	name := c.Query("name")
	age := c.DefaultQuery("age", "18")
	c.String(http.StatusOK, "获取到查询参数:name="+name+",age="+age)
}

// 路径参数,restful风格,路径参数可以和查询参数混合使用，并且值可以用,隔开，代表是一个数组
func PathGet(c *gin.Context) {
	// 路径参数
	// 127.0.0.1:8080/student/张三/18/其他

	name := c.Param("name")         // name为张三
	ageother := c.Param("ageother") // ageother为 /18/其他
	//使用*号可以获取到后面所有路径参数。，所以放在最后面，并且会包含/
	c.String(http.StatusOK, "获取到路径参数:name="+name+",ageother="+ageother)
}

// postform 表单获取参数
func PostForm(c *gin.Context) {
	// postform
	// 127.0.0.1:8080/student 使用post方法，Content-Type:application/x-www-form-urlencoded请求

	username := c.PostForm("username")
	age := c.DefaultPostForm("age", "18")
	c.String(http.StatusOK, "获取到表单参数:username="+username+",age="+age)
}

// 获取json参数
func JsonPost(c *gin.Context) {
	// json
	// 127.0.0.1:8080/student 使用post方法，Content-Type:application/json请求
	var stu struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}
	byteSlice, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "读取请求体失败")
		return
	}
	err = json.Unmarshal(byteSlice, &stu)
	if err != nil {
		c.String(http.StatusOK, "解析json失败")
		return
	}
	c.String(http.StatusOK, "获取到json参数:username="+stu.Username+",age="+strconv.Itoa(stu.Age))

}

// 上传单个文件，保存到服务器
func UploadSingleFile(c *gin.Context) {
	// 上传文件
	// 127.0.0.1:8080/student 使用post方法，Content-Type:multipart/form-data请求
	slog.Info("上传文件", "header", c.Request.Header)
	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("获取文件失败", "err", err)
		c.String(http.StatusOK, "获取文件失败")
		return
	}
	// 保存文件到服务器
	err = c.SaveUploadedFile(file, "./data/"+file.Filename)
	if err != nil {
		slog.Error("保存文件失败", "err", err)
		c.String(http.StatusOK, "保存文件失败")
		return
	}
	// 返回文件路径
	c.String(http.StatusOK, "上传成功:filename="+file.Filename)
	slog.Info("上传成功", "filename", file.Filename)

}

// 上传多个文件，保存到服务器
func UploadMultiFile(c *gin.Context) {
	// 上传文件
	// 127.0.0.1:8080/student 使用post方法，Content-Type:multipart/form-data请求
	slog.Info("上传文件", "header", c.Request.Header)

	files, err := c.MultipartForm()
	if err != nil {
		slog.Error("获取文件失败", "err", err)
		c.String(http.StatusOK, "获取文件失败")
		return
	}
	// 保存文件到服务器
	for _, file := range files.File["files"] {
		err = c.SaveUploadedFile(file, "./datas/"+file.Filename)
		if err != nil {
			slog.Error("保存文件失败", "err", err)
			c.String(http.StatusOK, "保存文件失败")
			return
		}
	}
	// 返回文件路径
	c.String(http.StatusOK, "上传成功")
	slog.Info("上传成功", "filename", files.File["file"])
}

// 参数绑定
// 绑定结构体参数
type Student struct {
	// form绑定的是get查询参数和post表单参数
	// json绑定的是post json参数
	// uri绑定的是路径参数
	Name     string   `form:"username" json:"name" uri:"user" xml:"name" yaml:"name" `
	Age      int      `form:"age" json:"age" uri:"age"`
	Keywords []string `form:"keywords" json:"keywords" uri:"keywords"` // 可以绑定html的复选框checkbox
}

// 表单绑定参数
func FormBind(c *gin.Context) {
	var stu Student
	err := c.ShouldBind(&stu)
	if err != nil {
		c.String(http.StatusOK, "绑定参数失败")
		return
	}
	c.String(http.StatusOK, "获取到表单参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
	slog.Info("绑定参数成功", "stu", stu)
}

// json 绑定参数
func JsonBind(c *gin.Context) {
	var stu Student
	err := c.ShouldBindJSON(&stu)
	if err != nil {
		c.String(http.StatusOK, "绑定参数失败")
		return
	}
	c.String(http.StatusOK, "获取到json参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
	slog.Info("绑定参数成功", "stu", stu)
}

// uri 绑定参数
// 127.0.0.1:8080/student/张三/18/关键词1,关键词2
func UriBind(c *gin.Context) {
	var stu Student
	err := c.ShouldBindUri(&stu)
	if err != nil {
		c.String(http.StatusOK, "绑定参数失败")
		return
	}
	c.String(http.StatusOK, "获取到uri参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
	slog.Info("绑定参数成功", "stu", stu)
}

// 使用ShouldBindBodyWith函数可以绑定任意格式的参数，比如json、XML、YAML、protobuf等
func BodyBind(c *gin.Context) {
	var stu Student
	if err := c.ShouldBindBodyWith(&stu, binding.JSON); err == nil {
		c.String(http.StatusOK, "获取到body参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
		return
	}

	// 绑定XML参数
	if err := c.ShouldBindBodyWith(&stu, binding.XML); err == nil {
		c.String(http.StatusOK, "获取到body参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
		return
	}
	// 绑定YAML参数
	if err := c.ShouldBindBodyWith(&stu, binding.YAML); err == nil {
		c.String(http.StatusOK, "获取到body参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
		return
	}
	// 绑定protobuf参数
	if err := c.ShouldBindBodyWith(&stu, binding.ProtoBuf); err == nil {
		c.String(http.StatusOK, "获取到body参数:name="+stu.Name+",age="+strconv.Itoa(stu.Age)+",keywords="+strings.Join(stu.Keywords, ","))
		return
	}
}
func main() {
	r := gin.Default()
	r.GET("/student", QueryGet)
	r.GET("/student/:name/*ageother", PathGet)
	r.POST("/student", PostForm)
	r.POST("/student/json", JsonPost)
	// 设置表单上限8M，默认是32M
	r.MaxMultipartMemory = 8 << 20
	// 上传单个文件
	r.POST("/student/uploadSingleFile", UploadSingleFile)
	// 上传多个文件
	r.POST("/student/uploadMultiFile", UploadMultiFile)
	// 表单 绑定参数
	r.POST("/student/FormBind", FormBind)
	// json 绑定参数
	r.POST("/student/JsonBind", JsonBind)
	// uri 绑定路径参数,路径参数是前缀优先的贪婪匹配机制，这里需要info，不然会和/student/:name/*ageother路由冲突
	r.GET("/student/info/:user/:age/:keywords", UriBind)

	// 启动服务
	r.Run("127.0.0.1:8080")

}
