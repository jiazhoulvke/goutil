package goutil

import (
	"os"
	"syscall"
)

//Mkdir 创建目录
func Mkdir(path string) error {
	oldMask := syscall.Umask(0)
	err := os.Mkdir(path, os.ModePerm)
	syscall.Umask(oldMask)
	return err
}
