package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	host  = "57.180.139.206"
	Index = fmt.Sprintf("beast_deploy_server_%s", host)
)

type CHook struct{}

func NewCHook() *CHook {
	return &CHook{}
}

func (hook CHook) Fire(entry *logrus.Entry) error {
	entry.Data["index"] = Index
	//log, err := entry.Bytes()
	//if err != nil {
	//	fmt.Println("Fire entry Error :", err)
	//}

	return nil
}

func (hook CHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func LogInit() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{}) //json格式数据
	logrus.AddHook(NewCHook())
}
