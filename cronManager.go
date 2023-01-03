package gotool

import (
	"log"
	"sync"
	"time"

	"github.com/robfig/cron"
)

// TenMinutesTask - 每日排程
func TenMinutesTask() {

	mi := "0,10,20,30,40,50"
	ho := "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23"

	c := cron.New()
	c.AddFunc("0 "+mi+" "+ho+" * * *", func() {
		//	do something
	})
	c.Start()
	log.Printf("Initialize HourTask Success!\n")
	log.Printf("Initialize http://127.0.0.1:9090/ \n")

}

type noCopy struct{} //nolint:unused

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// Timer -
type Timer struct {
	nocopy *noCopy
	timer  *time.Ticker
	Actor  sync.Map
}

// Start -
func (m *Timer) Start(d time.Duration) {
	m.timer = time.NewTicker(d)
	go func() {
		for {
			select {
			case <-m.timer.C:
				f := func(key, value interface{}) bool {
					f := value.(func())
					go f()
					return true
				}
				m.Actor.Range(f)
			}
		}
	}()
}

// Stop -
func (m *Timer) Stop() {
	m.timer.Stop()
}

// AddFunc -
func (m *Timer) AddFunc(name string, act func()) bool {

	_, ok := m.Actor.Load(name)
	if !ok {
		log.Printf("開始一個任務 : %v", name)
		m.Actor.Store(name, act)
		return false
	}
	return true
}

// RemoveFunc -
func (m *Timer) RemoveFunc(name string) {

	_, ok := m.Actor.Load(name)
	if ok {
		log.Printf("結束一個任務 : %v", name)
		m.Actor.Delete(name)
	}
}
