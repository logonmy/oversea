package utils

import (
	"os"
	"time"
	"fmt"
)

func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)

	if err != nil {
		fmt.Print("------------%v", err)
		fmt.Println( os.Getwd())
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Print("--------2----%v", err)
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

func RealPath(filePath string) string {
	return os.ExpandEnv(filePath)
}