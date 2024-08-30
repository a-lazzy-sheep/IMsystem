package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// lowwer case
func Md5EncodeLower(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// upper case
func Md5EncodeUpper(str string) string {
	return strings.ToUpper(Md5EncodeLower(str))
}

// 
func MakePassword(plainpwd, salt string) string {
	return Md5EncodeLower(plainpwd + salt)
}

func CheckPassword(plainpwd, salt, hashedpwd string) bool {
	return MakePassword(plainpwd, salt) == hashedpwd
}