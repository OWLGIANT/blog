package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func auth(ip string) int {
	resp, err := http.Get(fmt.Sprintf("http://auth.thousandquant.com:8500/add?ip=%s", ip))
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func sleep(delay time.Duration) {
	time.Sleep(delay * time.Millisecond)
}

func PermissionAdd(ipList []string) {
	if len(ipList) <= 0 {
		logrus.Error("ip 不能为空")
		return
	}
	for i := 0; i < len(ipList); i++ {
		status := auth(ipList[i])
		fmt.Printf("ip %s %d\n", ipList[i], status)
		sleep(200)
	}
}
