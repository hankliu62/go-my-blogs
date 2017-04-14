package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5Encode create md5 string
func Md5Encode(str string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))

	cipherStr := md5Hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
