package goutil

import (
	"bytes"
	"math/rand"
	"os"
	"time"
)

const (
	randomStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(rand.Int63())
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
