package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5 ...
func MD5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}