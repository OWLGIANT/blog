package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

type Node struct {
	data int
	next *Node
}

func allocateMemory() {
	for i := 0; i < 100000; i++ {
		s := make([]byte, 1024) // 模拟内存分配
		_ = s
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 开始内存分析
	go startMemoryProfiling()

	// 创建一个头节点
	head := &Node{data: 0, next: nil}
	current := head

	// 模拟创建大量节点，但不释放它们
	for i := 1; i < 10000; i++ {
		newNode := &Node{data: i, next: nil}
		current.next = newNode
		current = newNode
	}
	println("===============================")

	// 模拟内存泄漏
	for {
		allocateMemory()
		time.Sleep(1 * time.Second)
	}
}
func startMemoryProfiling() {
	// 每隔一段时间收集一次内存分析信息
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			memProfile, err := os.Create("mem.pprof")
			if err != nil {
				fmt.Println("Error creating memory profile: ", err)
				return
			}
			// 收集内存分析信息
			pprof.WriteHeapProfile(memProfile)
			memProfile.Close()
		}
	}
}
