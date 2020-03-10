package series

import (
	"errors"
	"fmt"
)

// Fibonacci数列相关

// init s
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 定义全局错误变量
var LessThan2 = errors.New("Less Than 2")
var OverThan100 = errors.New("Over Than 100")

// 获取Fibonacci数列
func GetFibonacciSeries(len int) ([]int, error) {
	if len < 2 {
		return nil, LessThan2
	}

	if len > 100 {
		return nil, OverThan100
	}

	ret := []int{1, 1}
	for i := 2; i < len; i++ {
		ret = append(ret, ret[i-2] + ret[i-1])
	}
	return ret, nil
}
