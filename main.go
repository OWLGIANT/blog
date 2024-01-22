package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := "my_file.txt"

	// 打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建文件锁
	lock := flock.NewFlock(filePath)

	// 尝试加锁
	locked, err := lock.TryLock()
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return
	}

	if !locked {
		fmt.Println("Another process holds the lock.")
		return
	}

	fmt.Println("Acquired lock")

	// 模拟业务逻辑
	time.Sleep(10 * time.Second)

	// 释放锁
	err = lock.Unlock()
	if err != nil {
		fmt.Println("Error releasing lock:", err)
		return
	}

	fmt.Println("Released lock")
}
