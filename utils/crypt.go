package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
	"strings"
)

// MD5 ...
func MD5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func NewUUID() string {
	v4, _ := uuid.NewV4()
	return uuid.NewV5(v4, strconv.FormatInt(time.Now().UnixNano(), 10)).String()
}
func NewNoDashUUID() string {
	return strings.Replace(NewUUID(), "-", "", -1)
}