package util

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密
// 对文件名进行格式化：对文件名用MD5加密后在进行写入，避免直接暴露原始名称
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
