package account

import (
	"os"
	"path/filepath"
	"src/global"
)

// 创建存放头像的文件夹
func InitAvatarFolder() {
	// TODO:每次编译前需要修改为CurrentDir，debug时用ParentDir
	global.AvatarsDir = filepath.Join(global.CurrentDir, "/avatars")
	// 如果没找到avatars，则创建avatars文件夹
	if _, err := os.Stat(global.AvatarsDir); os.IsNotExist(err) {
		os.Mkdir(global.AvatarsDir, 0755)
	}
}
