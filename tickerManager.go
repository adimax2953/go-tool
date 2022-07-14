package gotool

import (
	"sync"
	"time"

	LogTool "github.com/adimax2953/log-tool"
)

//Ticker 定義一個Ticker,用於時間控制
type Ticker struct {
	id    uint64
	cb    func()
	timer *time.Timer
	stop  chan int
}

//TickerManager 封裝一個tick的管理類,避免每次都要寫一大堆
type TickerManager struct {
	autoid    uint64
	mapTicker map[uint64]*Ticker
	mux       sync.Mutex
}

//Init 初始化tickermanager, 就是設置起始id
func (tm *TickerManager) Init() {
	LogTool.LogSystem("init TickerManager")
	tm.autoid = 0
	tm.mapTicker = make(map[uint64]*Ticker)
}

//DelayExec 延遲執行函數
func (tm *TickerManager) DelayExec(callback func(), d time.Duration) uint64 {
	tm.mux.Lock()
	tm.autoid++
	tm.mux.Unlock()

	go func() {
		newtimer := time.NewTimer(d)
		ticker := &Ticker{
			id:    tm.autoid,
			cb:    callback,
			timer: newtimer,
			stop:  make(chan int),
		}
		tm.mapTicker[tm.autoid] = ticker

		select {
		case <-newtimer.C:
			callback()
		case <-ticker.stop:
			ticker.timer.Stop()
		}

		delete(tm.mapTicker, ticker.id)
	}()

	return tm.autoid
}

//ClearTicker 清除tick, 停止tick
func (tm *TickerManager) ClearTicker(id uint64) {
	ticker, ok := tm.mapTicker[id]
	if ok {
		ticker.stop <- 1
	}
}

//GetTickManager 得到tickmanage
func GetTickManager() *TickerManager {
	return &tickerManager
}

var tickerManager TickerManager

func init() {
	tickerManager.Init()
}
