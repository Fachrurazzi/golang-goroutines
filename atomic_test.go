package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				x = x + 1
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
