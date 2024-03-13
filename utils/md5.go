package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// @将字符串加密成 md5   md5 校验主要是校验数据在传输的过程中有没有产生错误
func String2md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
