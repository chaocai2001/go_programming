package type_def

import (
	"fmt"
	"time"
)

type IntConvertionFn func(n int) int

func timeSpent(inner IntConvertionFn) IntConvertionFn {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time escaped ", time.Since(start).Seconds())
		return ret
	}
}
