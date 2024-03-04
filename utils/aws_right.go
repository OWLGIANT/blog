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

func auth2(ip string) int {
	resp, err := http.Get(fmt.Sprintf("http://server.thousandquant.com:8500/add?ip=%s", ip))
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
		status2 := auth2(ipList[i])
		fmt.Printf("ip %s %d %d\n", ipList[i], status, status2)
		sleep(200)
	}
}
