package main

import (
	"os"
	"path/filepath"
)

// 创建存放头像的文件夹
func initAvatarFolder() {
	// TODO:每次编译前需要修改为currentDir，debug时用parentDir
	avatarsDir = filepath.Join(parentDir, "/avatars")
	// 如果没找到avatars，则创建avatars文件夹
	if _, err := os.Stat(avatarsDir); os.IsNotExist(err) {
		os.Mkdir(avatarsDir, 0755)
	}
}

// TODO:压缩头像分辨率为512*512
// func compressAvatar(filePath string) {
// 	var width uint = 64
// 	var height uint = 64
// }

// TODO:接入头像检测API
