package gopackage

import (
	"common/series"
	"testing"
)
// package 包
// 基本复用模块
// 以首字母大写来表明可被包外代码访问
// 代码的package可以和所在的目录可以不一致
// 同一个目录里的Go代码的package要保持一致

// proj gopath E:\src\go_learning

func TestPackage(t *testing.T) {
	for i := -1; i < 5; i ++ {
		if fibonacci, err := series.GetFibonacciSeries(i); err == nil {
			t.Log(len(fibonacci))
		} else {
			t.Error(err)
		}
	}

	for i := 99; i < 102; i ++ {
		if fibonacci, err := series.GetFibonacciSeries(i); err == nil {
			t.Log(len(fibonacci))
		} else {
			t.Error(err)
		}
	}
}