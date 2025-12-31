package io

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 存储所有遍历到的路径
var allPaths []string

// 遍历指定目录下所有JSON文件，返回文件路径列表和错误
func walkJSONFiles(rootDir string) ([]string, error) {
	// 校验目录是否存在
	info, err := os.Stat(rootDir)
	if err != nil {
		return nil, fmt.Errorf("目录不存在或无法访问: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("指定路径不是目录: %s", rootDir)
	}

	var jsonFiles []string
	// 计算遍历次数
	traverseCount := 0

	// 递归遍历目录,第一个路径是根目录
	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		// 处理遍历过程中的错误（如权限不足）
		if err != nil {
			fmt.Printf("遍历文件失败: %s, 错误: %v\n", path, err)
			return nil // 忽略错误继续遍历其他文件
		}
		traverseCount++

		// 记录根目录
		allPaths = append(allPaths, path)

		// // 输出当前遍历路径（可选，用于调试）
		// fmt.Printf("当前遍历路径: %s\n", path)

		// // 跳过目录（只处理文件）
		// if d.IsDir() {
		// 	// 输出当前遍历目录（可选，用于调试）
		// 	fmt.Printf("当前遍历目录名称: %s\n", d.Name())
		// 	return nil
		// }

		// 获取文件扩展名（转小写，兼容.JSON/.Json等情况）
		ext := strings.ToLower(filepath.Ext(path))
		// 判断是否为JSON文件
		if ext == ".json" {
			jsonFiles = append(jsonFiles, path)
		}

		return nil
	})

	// 输出遍历次数（可选，用于调试）
	fmt.Printf("共遍历 %d 个文件/目录\n", traverseCount)

	if err != nil {
		return nil, fmt.Errorf("遍历目录失败: %w", err)
	}

	return jsonFiles, nil
}

func Traverse() {
	// 指定要遍历的根目录（可替换为自己的目录路径，如 "./data"）
	rootDir := `D:\data\项目\1126多模态方案\toolchain\1205多模互动视频\label_predict`

	// 遍历JSON文件
	jsonFiles, err := walkJSONFiles(rootDir)
	if err != nil {
		fmt.Printf("遍历JSON文件失败: %v\n", err)
		os.Exit(1)
	}

	// 输出结果
	if len(jsonFiles) == 0 {
		fmt.Printf("目录 %s 下未找到JSON文件\n", rootDir)
		return
	}

	fmt.Printf("共找到 %d 个JSON文件：\n", len(jsonFiles))
	// for _, file := range jsonFiles {
	// 	fmt.Println("-", file)
	// }
	const N = 5

	for i := 0; i < N; i++ {
		fmt.Println(allPaths[i])
	}
}
