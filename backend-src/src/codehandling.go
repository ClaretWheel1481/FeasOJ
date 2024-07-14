package main

import (
	"os"
	"path/filepath"
)

// 初始化文件夹
func initCodeFolder() {
	codeDir = filepath.Join(parentDir, "codefiles")
	// 如果没找到codefiles，则创建codefiles文件夹
	if _, err := os.Stat(codeDir); os.IsNotExist(err) {
		os.Mkdir(codeDir, 0755)
	}
}
