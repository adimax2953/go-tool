package worktool

import (
	"github.com/panjf2000/ants/v2"
)

type WorkPool struct {
	Pool *ants.Pool
}

// NewWorkPool - 建立一個工作池 poolsize：要有多少線呈 return WorkPool
func NewWorkPool(poolsize int) *WorkPool {
	pool, _ := ants.NewPool(poolsize, ants.WithPreAlloc(true))
	return &WorkPool{Pool: pool}
}

// NewWorkPool - 建立WorkPool poolsize：要有多少線呈 return PoolWithFunc
func (wp *WorkPool) NewWorkPoolWithFunc(poolsize int, poolfunc func(interface{})) (*ants.PoolWithFunc, error) {
	pf, err := ants.NewPoolWithFunc(poolsize, poolfunc, ants.WithPreAlloc(true))
	return pf, err
}

// ChangePoolSize - 修改一個工作池 poolsize：要改多少線呈
func (wp *WorkPool) ChangePoolSize(poolsize int) {
	wp.Pool.Tune(poolsize)
}

// SubmitTask - 提交一個工作到WorkPool task : 要做事的funtion
func (wp *WorkPool) SubmitTask(task func()) error {
	return wp.Pool.Submit(task)
}

// Release - 釋放WorkPool
func (wp *WorkPool) Release() {
	wp.Pool.Release()
}

// Reboot - 重啟WorkPool
func (wp *WorkPool) Reboot() {
	wp.Pool.Reboot()
}
