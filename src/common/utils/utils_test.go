package utils

import (
	"testing"
	"time"
)

func slowFunc(op int) int {
	time.Sleep(time.Second * 1) // 1s
	return op
}

func TestCalcRuntime(t *testing.T) {
	f := CalcRuntime(slowFunc)
	f(1)
}
