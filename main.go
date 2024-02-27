package main

import (
	"blog/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	_ "net/http/pprof"
	"time"
)

func main() {
	utils.LogInit()

	logrus.Info("====================日志插入测试===============================")

	fmt.Println("============defedwefewfewfefefef================")
	time.Sleep(time.Minute)
}
