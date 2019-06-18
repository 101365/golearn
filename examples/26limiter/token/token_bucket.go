package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type TokenBucket struct {
	capacity int64 // 容量
	tokens   int64 // 当前token数量
	sync.Mutex
}

// 创建
func NewTokenBucket(cap int64, rate time.Duration) *TokenBucket {

	t := &TokenBucket{
		capacity: cap,
		tokens:   0,
	}

	ticker := time.NewTicker(rate)

	// 定时向桶中添加令牌
	go func() {
		for range ticker.C {
			if atomic.LoadInt64(&t.tokens) >= t.capacity {
				continue
			}
			atomic.AddInt64(&t.tokens, 1)
		}
	}()

	return t
}

// 获取
func (b *TokenBucket) Get(n int64) bool {

	b.Lock()
	defer b.Unlock()

	if atomic.LoadInt64(&b.tokens) < n {
		return false
	}

	atomic.AddInt64(&b.tokens, -n)

	return true
}

func (b *TokenBucket) Display() {
	fmt.Println("capacity=", b.capacity, ", tokens=", b.tokens)
}

func main() {
	var wg sync.WaitGroup

	// 容量3 每秒最多2个令牌
	limiter := NewTokenBucket(3, time.Millisecond*500)

	time.Sleep(time.Second * 3)

	for i := 0; i < 200; i++ {
		wg.Add(1)

		//fmt.Println("Create req", i, time.Now())
		go func(i int) {
			if limiter.Get(1) {
				limiter.Display()
				fmt.Println("Respon req", i, time.Now())
			}
			wg.Done()
		}(i)

		time.Sleep(20 * time.Millisecond)
	}
	wg.Wait()
}
