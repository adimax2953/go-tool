package kafkatool

import (
	"context"
	LogTool "github.com/adimax2953/log-tool"
	"golang.org/x/time/rate"
	"time"
)

type WorkerCircuitBreaker struct {
	consumeTokenPerRequest int // 每次消耗令牌量
	maxDataSize            int // 滿足此條件表示要消耗令牌
	limiter                *rate.Limiter
}

// NewWorkerCircuitBreaker
// bucketVolume: 令牌桶容量
// consumeTokenPerRequest: 每次消耗令牌量
// maxDataSize: 滿足此條件表示要消耗令牌
// e.g. bucketVolume = 10, consumeTokenPerRequest = 3, maxDataSize = 100
// 表示每秒產生1個token, 但每次消耗3個token, 當資料筆數大於等於100時, 消耗三個token
// 當桶子內的token 不夠時, 會等待下一個token產生
// 用這方式降低資料庫寫入壓力
func NewWorkerCircuitBreaker(bucketVolume, consumeTokenPerRequest, maxDataSize int) *WorkerCircuitBreaker {
	return &WorkerCircuitBreaker{
		limiter:                rate.NewLimiter(rate.Every(time.Second), bucketVolume),
		consumeTokenPerRequest: consumeTokenPerRequest,
		maxDataSize:            maxDataSize,
	}
}

func (r *WorkerCircuitBreaker) Check(ctx context.Context, dataCount int) {
	if dataCount < r.maxDataSize {
		return
	}
	start := time.Now()
	err := r.limiter.WaitN(ctx, r.consumeTokenPerRequest)
	if err != nil {
		LogTool.LogErrorf("Check", "limiter error: %v", err)
		return
	}
	LogTool.LogInfof("Check", "circuit breaker time: %v", time.Since(start))
}
func (r *WorkerCircuitBreaker) CheckContinuouslyWriting(ctx context.Context) {
	err := r.limiter.WaitN(ctx, r.consumeTokenPerRequest)
	if err != nil {
		LogTool.LogErrorf("Check", "limiter error: %v", err)
		return
	}
}
