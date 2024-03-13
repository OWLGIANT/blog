package utils

import "fmt"

// 分片处理
func RobotBatchCreatRequest(robots []string, batchSize int) {
	numData := len(robots)
	numBatches := (numData + batchSize - 1) / batchSize // 计算批次数，向上取整
	for i := 0; i < numBatches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize
		fmt.Println(end, numData)
		if end > numData {
			end = numData

		}

		fmt.Println(robots[start:end])

	}
}
