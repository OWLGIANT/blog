package utils

import "sync"

type LimitWaitGroup struct {
	sig       chan struct{}
	wg        sync.WaitGroup
	PublicVar sync.Map
}

func NewGoroutineServer(limit int) *LimitWaitGroup {
	return &LimitWaitGroup{
		sig:       make(chan struct{}, limit),
		wg:        sync.WaitGroup{},
		PublicVar: sync.Map{},
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
