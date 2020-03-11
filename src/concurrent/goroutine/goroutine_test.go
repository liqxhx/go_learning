package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 协程

// System Thread
// Processor go实现的协程处理器
// goroutine

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go 在协程中运行
		go func(op int) {
			fmt.Println(op)
		}(i)

		//go func() {
		//	fmt.Println(i)
		//}()
	}

	time.Sleep(time.Second * 2)
}

func TestCounterNonConcurrent(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 2)
	t.Log(counter)
}

func TestCounter(t *testing.T) {
	counter := 0

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				mutex.Unlock()
			}()
			mutex.Lock()
			counter++
		}()
	}

	wg.Wait() //
	t.Log(counter)
}
