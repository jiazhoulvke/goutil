package goutil

import (
	"bytes"
	"math/rand"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

const (
	randomStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Int63n(9999999999)) * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
}

//RandomString 生成随机字符串
func RandomString(length int, randomStrs ...string) string {
	var rs string
	if len(randomStrs) > 0 && len(randomStrs[0]) > 0 {
		rs = randomStrs[0]
	} else {
		rs = randomStr
	}
	b := bytes.NewBuffer([]byte{})
	rsLen := len(rs)
	for i := 0; i < length; i++ {
		b.WriteByte(byte(rs[rand.Intn(rsLen)]))
	}
	return b.String()
}

//IsExist 文件或目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//Mkdir 创建目录
func Mkdir(path string) error {
	oldMask := syscall.Umask(0)
	err := os.Mkdir(path, os.ModePerm)
	syscall.Umask(oldMask)
	return err
}

//CreateParentDir 递归创建上级目录
func CreateParentDir(p string) error {
	if p == "." {
		return nil
	}
	if IsExist(p) {
		return nil
	}
	parentDir := filepath.Dir(p)
	if !IsExist(parentDir) {
		if err := CreateParentDir(parentDir); err != nil {
			return err
		}
	}
	return Mkdir(p)
}
