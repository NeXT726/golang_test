//学习go关键字，创建goroutine。判断创建的协程间的运行顺序。

package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go println(i)
	}
	runtime.Gosched()
	time.Sleep(time.Second)
}
