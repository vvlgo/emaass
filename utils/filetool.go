package utils

import (
	"os"
	"strings"
)

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取绝对路径
func getPath() string {
	dir, _ := os.Getwd()

	return dir + "/"
}

//获取上级路径
func getParentDirectory() string {
	dir, _ := os.Getwd()
	index := strings.LastIndex(dir, "/")
	runes := []rune(dir)
	return string(runes[:index+1])
}
