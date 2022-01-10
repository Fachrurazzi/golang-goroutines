package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunsAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hellow")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunsAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
