package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	host  = "114.55.134.26"
	Index = fmt.Sprintf("public_server_%s", host)
)

type CHook struct{}

func NewCHook() *CHook {
	return &CHook{}
}

func (hook CHook) Fire(entry *logrus.Entry) error {
	entry.Data["index"] = Index
	log, err := entry.Bytes()
	if err != nil {
		fmt.Println("Fire entry Error :", err)
	}
	response, err := HTTPInsert(http.MethodPost, fmt.Sprintf("http://%s:8080/logrus/insert", host), log)
	if err != nil {
		fmt.Println("HTTPRequest Error :", err)
	}
	//var body struct {
	//	Code int                     `json:"code"`
	//	Msg  string                  `json:"msg"`
	//}
	//if err = json.Unmarshal(response, &body); err != nil {
	//	logrus.Error(err)
	//}

	fmt.Println("========response==========", string(response))

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
