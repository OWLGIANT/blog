package utils

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func StartMemoryProfiling(pid int) {
	logrus.Infof("每隔一段时间收集一次内存分析信息,内存分析信息写入")
	var (
		maxSys   uint64
		maxAlloc uint64
	)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			memProfile, err := os.Create(fmt.Sprintf("WatchMem%v.pprof", pid))
			if err != nil {
				logrus.Error("Error creating memory profile: ", err)
				return
			}
			// 收集内存分析信息
			pprof.WriteHeapProfile(memProfile)
			memProfile.Close()

			MemStatsProfile, err := os.Create(fmt.Sprintf("WatchMemStats%v.txt", pid))
			if err != nil {
				logrus.Error("Error creating MemStats Profile: ", err)
				return
			}
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			if m.Alloc/1024 > maxAlloc {
				maxAlloc = m.Alloc / 1024
			}
			if m.Sys/1024 > maxSys {
				maxSys = m.Sys / 1024
			}
			// 创建一个writer
			writer := bufio.NewWriter(MemStatsProfile)
			content := fmt.Sprintf(" maxSys:%vKB\n Sys:%vKB\n MaxAlloc:%vKB\n Alloc:%vKB\n TotalAlloc:%vKB\n NumGC:%v次\n", maxSys, m.Sys/1024, maxAlloc, m.Alloc/1024, m.TotalAlloc/1024, m.NumGC)
			_, err = writer.WriteString(content)
			if err != nil {
				fmt.Printf("write file err, %v", err)
			}
			writer.Flush()
			MemStatsProfile.Close()
		}
	}
}

func MeMLeak() {
	resultString := ""
	for {
		resultString = resultString + "======================天上天下 , 唯朕独尊======================"
	}
}
