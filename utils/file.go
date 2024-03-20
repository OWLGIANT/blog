package utils

import (
	"fmt"
	"os"
)

func WriteFile(content, filePath string) {
	// 打开文件以供写入，如果文件不存在则创建文件，如果文件已存在则会覆盖文件内容
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}
	fmt.Println("写入文件成功")
}
