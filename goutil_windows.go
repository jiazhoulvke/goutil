package goutil

import (
	"os"
)

//Mkdir 创建目录
func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}
