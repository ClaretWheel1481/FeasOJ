package main

import (
	"os"
	"path/filepath"
)

var avatarsDir string

// 创建存放头像的文件夹
func initAvatarFolder() {
	avatarsDir = filepath.Join(parentDir, "avatars")
	// 如果没找到avatars，则创建avatars文件夹
	if _, err := os.Stat(avatarsDir); os.IsNotExist(err) {
		os.Mkdir(configsDir, 0755)
	}
}

// TODO:压缩文件分辨率为100*100
