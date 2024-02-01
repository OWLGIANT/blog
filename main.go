package main

import (
	"blog/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

type LimitWaitGroup struct {
	sig       chan struct{}
	wg        sync.WaitGroup
	publicVar sync.Map
}

func NewGoroutineServer(limit int) *LimitWaitGroup {
	return &LimitWaitGroup{
		sig:       make(chan struct{}, limit),
		wg:        sync.WaitGroup{},
		publicVar: sync.Map{},
	}
}

func (s *LimitWaitGroup) Add() {
	s.sig <- struct{}{}
	s.wg.Add(1)
}

func (s *LimitWaitGroup) Down() {
	<-s.sig
	s.wg.Done()
}

func (s *LimitWaitGroup) Wait() {
	s.wg.Wait()
}

func main() {

	outPut, err := utils.SSHCmd("114.55.134.26", "pwd;pwd")
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(strings.Split(outPut, "\n"))
}

func Check() {
	wg := NewGoroutineServer(2)
	ips := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}
	var (
		strategy = "strategy"
		version  = "version"
	)

	for _, v := range ips {
		wg.Add()
		go wg.Logic(v, strategy, version)
	}
	wg.Wait()
	details := fmt.Sprintf("如下服务器策略 %v 发布失败\n", "strategy")
	wg.publicVar.Range(func(key, dif any) bool {
		ip, ok := key.(string)
		if ok {
			details += ip + "\n"
		}
		return true
	})
	fmt.Println(details)
}

func (s *LimitWaitGroup) Logic(ip, strategy, version string) {
	defer s.Down()
	fmt.Println("========ip, strategy, version========", ip, strategy, version)
	s.publicVar.Store(ip, true)

	//if !CheckLastVersion(ip, strategy, version){
	//
	//}
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
