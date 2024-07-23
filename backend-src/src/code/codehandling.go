package code

import (
	"os"
	"path/filepath"
	"src/global"
)

// 初始化文件夹
func InitCodeFolder() {
	// TODO:每次编译前需要修改为currentDir，debug时用parentDir
	global.CodeDir = filepath.Join(global.ParentDir, "/codefiles")
	// 如果没找到codefiles，则创建codefiles文件夹
	if _, err := os.Stat(global.CodeDir); os.IsNotExist(err) {
		os.Mkdir(global.CodeDir, 0755)
	}
}
