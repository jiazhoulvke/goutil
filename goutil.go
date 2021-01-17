package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	randomStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//RandomString 生成随机字符串
func RandomString(length int, randomStrs ...string) string {
	rand.Seed(time.Now().UnixNano())
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

//JSON 返回对象的JSON字符串
func JSON(o interface{}) string {
	b := bytes.NewBufferString("")
	jsonEncoder := json.NewEncoder(b)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(o)
	return b.String()
}

//InStringSlice 判断字符串是否在列表中
func InStringSlice(v string, s []string) bool {
	for i := 0; i < len(s); i++ {
		if v == s[i] {
			return true
		}
	}
	return false
}

//InIntSlice 判断数字是否在列表中
func InIntSlice(v int, s []int) bool {
	for i := 0; i < len(s); i++ {
		if v == s[i] {
			return true
		}
	}
	return false
}

//InInt64Slice 判断数字是否在列表中
func InInt64Slice(v int64, s []int64) bool {
	for i := 0; i < len(s); i++ {
		if v == s[i] {
			return true
		}
	}
	return false
}

//JoinIntSlice 类似strings.Join
func JoinIntSlice(s []int, sep string) string {
	ss := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		ss = append(ss, strconv.Itoa(s[i]))
	}
	return strings.Join(ss, sep)
}

//JoinInt64Slice 类似strings.Join
func JoinInt64Slice(s []int64, sep string) string {
	ss := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		ss = append(ss, strconv.FormatInt(s[i], 10))
	}
	return strings.Join(ss, sep)
}

//String2IntSlice 字符串转int列表
func String2IntSlice(str string, sep string) ([]int, error) {
	sSlice := strings.Split(str, sep)
	iSlice := make([]int, 0)
	for _, s := range sSlice {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return iSlice, err
		}
		iSlice = append(iSlice, n)
	}
	return iSlice, nil
}

//String2Int64Slice 字符串转int64列表
func String2Int64Slice(str string, sep string) ([]int64, error) {
	sSlice := strings.Split(str, sep)
	iSlice := make([]int64, 0)
	for _, s := range sSlice {
		if s == "" {
			continue
		}
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return iSlice, err
		}
		iSlice = append(iSlice, n)
	}
	return iSlice, nil
}

//FindInSet 类似mysql的FIND_IN_SET
func FindInSet(v interface{}, str string) bool {
	o := fmt.Sprint(v)
	l := strings.Split(str, ",")
	for _, s := range l {
		if s == "" {
			continue
		}
		if o == s {
			return true
		}
	}
	return false
}

//RunesFilter 过滤字符
// RunesFilter("abc,def!ghi",",!") == "abcdefghi"
func RunesFilter(s string, cutset string) string {
	if s == "" || cutset == "" {
		return s
	}
	cutsetRute := []rune(cutset)
	b := bytes.NewBuffer([]byte(""))
FORLOOP:
	for _, r := range []rune(s) {
		for _, c := range cutsetRute {
			if c == r {
				continue FORLOOP
			}
		}
		b.WriteRune(r)
	}
	return b.String()
}

//IsZipFile 判断是不是zip文件
func IsZipFile(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()
	b := make([]byte, 4, 4)
	n, err := f.Read(b)
	if err != nil {
		return false, err
	}
	if n != 4 {
		return false, nil
	}
	return IsZipHead(b), nil
}

//IsZipHead 判断是不是zip文件头
func IsZipHead(b []byte) bool {
	if len(b) < 4 {
		return false
	}
	return b[0] == 0x50 && b[1] == 0x4b && b[2] == 0x03 && b[3] == 0x04
}
