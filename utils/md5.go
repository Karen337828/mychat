package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MakePassword 加密
func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd + salt)
}

// ValidPassword 校验密码是否正确
func ValidPassword(plainpwd, salt string, password string) bool {
	md := MD5Encode(plainpwd + salt)
	fmt.Println(md + "--------" + password)
	return md == password
}
