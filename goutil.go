package goutil

import (
	"bytes"
	"math/rand"
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
