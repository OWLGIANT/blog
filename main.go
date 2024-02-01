package main

import (
	"blog/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {

	//
}

func CheckLastVersion(ip, strategy, version string) bool {
	command := fmt.Sprintf(`cd /root/supervisor;./%s -v  | grep -o '"BuildTime":"[^"]*"' | awk -F'"' '{print $4}'`, strategy)
	outPut, err := utils.SSHCmd(ip, command)
	if err != nil {
		logrus.Error(err)
		return false
	}
	if outPut != version {
		return false
	}
	return true
}
