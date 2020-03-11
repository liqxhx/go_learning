package utils

import (
	"fmt"
	"time"
)

type FUN_INT_INT func(int) int

func CalcRuntime(in FUN_INT_INT) FUN_INT_INT {
	return func(op int) int {
		start := time.Now()
		ret := in(op)
		fmt.Println(&in, "TimeSpent:", time.Since(start).Seconds())
		return ret
	}
}
