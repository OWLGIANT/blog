package utils

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func Compress(csvFilePath string) bool {
	csvFilePath = "doc/cmd/" + csvFilePath
	// 压缩后的 Gzip 文件路径
	gzipFilePath := fmt.Sprintf("%v.gz", csvFilePath)

	// 打开 CSV 文件
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		logrus.Warnf("[compress] Error opening CSV file: %v", err)
		return false
	}
	defer csvFile.Close()

	// 创建 Gzip 文件
	gzipFile, err := os.Create(gzipFilePath)
	if err != nil {
		logrus.Warnf("[compress] Error creating Gzip file: %v", err)
		return false
	}
	defer gzipFile.Close()

	// 创建 Gzip Writer
	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	// 创建 CSV Reader
	csvReader := csv.NewReader(csvFile)

	// 创建 CSV Writer
	csvWriter := csv.NewWriter(gzipWriter)
	defer csvWriter.Flush()

	// 逐行读取 CSV 文件并写入 Gzip 文件
	for {
		record, err := csvReader.Read()
		if err != nil {
			break // 读取完成或出错时退出循环
		}
		if err := csvWriter.Write(record); err != nil {
			logrus.Warnf("Error writing to Gzip file: %v", err)
			return false
		}
	}

	logrus.Warnf("CSV file compressed to Gzip successfully.")
	return true
}
