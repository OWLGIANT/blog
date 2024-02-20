package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("http://localhost:6060/debug/pprof")
		http.ListenAndServe("localhost:6060", nil)
	}()
	// 手动触发 Heap Dump
	f, err := os.Create("heapdump.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("===开始=========", time.Now().Format(time.DateTime))

	for i := 0; i < 100000; i++ {
		var data []string
		data = append(data, "test")
	}
	fmt.Println("===结束=========", time.Now().Format(time.DateTime))
	runtime.GC() // 执行垃圾回收
	pprof.WriteHeapProfile(f)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Hour)
}

func Write() {
	var data []string
	data = append(data, "test")
}
