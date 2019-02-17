package models

import (
	"time"
)

type dbMaintainer interface {
	AddFunc(func(), time.Duration)
	Stop()
}

type DBMaintainer struct {
	MaintainFuncs []func()
	stop chan int
}

func (self * DBMaintainer) AddFunc(MaintainFunc func(), duration time.Duration){
	tf := self.timingFunc(MaintainFunc, duration)
	self.MaintainFuncs = append(self.MaintainFuncs, tf)
	go tf()
}

func (self *DBMaintainer)Stop() {
	for i := 0; i < len(self.MaintainFuncs); i++{
		self.stop <- 1
	}
}

func (self *DBMaintainer)timingFunc(f func(), duration time.Duration) func(){
	ticker := time.NewTicker(duration)
	return func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-self.stop:
				ticker.Stop()
				return
			}
		}
	}
}