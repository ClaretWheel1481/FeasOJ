package main

import (
	"os"
	"path/filepath"
)

// 创建存放头像的文件夹
func initAvatarFolder() {
	avatarsDir = filepath.Join(parentDir, "avatars")
	// 如果没找到avatars，则创建avatars文件夹
	if _, err := os.Stat(avatarsDir); os.IsNotExist(err) {
		os.Mkdir(configsDir, 0755)
	}
}

// TODO:压缩头像分辨率为512*512
// func compressAvatar(filePath string) {
// 	var width uint = 64
// 	var height uint = 64
// }

// TODO:接入AI模型对上传的头像进行检测
