package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("http://localhost:6060/debug/pprof")
		http.ListenAndServe("localhost:6060", nil)
	}()
	fmt.Println(time.Now().Format(time.DateTime))
	tickInfo := time.NewTicker(time.Second)
	defer tickInfo.Stop()
	for {
		select {
		case <-tickInfo.C:
			//Write()
			var data []string
			data = append(data, "test")
		}
	}
}

func Write() {
	var data []string
	data = append(data, "test")
}
