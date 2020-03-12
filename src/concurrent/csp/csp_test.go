package csp

import (
	"fmt"
	"testing"
	"time"
)

// 模拟耗时业务
func service() string {
	time.Sleep(time.Millisecond * 50) // 50 毫秒
	return "Done"
}

// 另一个任务
func otherTask() {
	fmt.Println("Starting Other Task...")
	time.Sleep(time.Millisecond * 100) // 100 毫秒
	fmt.Println("Stop Other Task.")
}

// 串行
func TestService(t *testing.T) {
	t.Log(service())
	otherTask()
}

// 对service的异步封装
func asyncService() chan string {
	retCh := make(chan string)
	go func() {
		fmt.Println("service() invoke")
		retStr := service()
		fmt.Println("service() Return result")
		retCh <- retStr
		fmt.Println("service() exit")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	ch := asyncService()
	// str := <-ch
	otherTask()
	str := <-ch
	t.Log("!!!!!", str)
}
