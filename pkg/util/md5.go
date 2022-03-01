package util

import (
	"crypto/md5"
	"encoding/hex"
)

//EnCodeMD5 formatting file name using hash and md5 encoder
func EnCodeMD5(value string) string {
	//md5 instance
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}


