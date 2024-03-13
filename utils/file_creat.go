package utils

import (
	"os"
	"path/filepath"
)

func OsCreate(filePath string) (file *os.File, err error) {
	//创建文件夹
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return
	}
	//创建文件
	file, err = os.Create(filePath)
	if err != nil {
		return
	}
	return
}
