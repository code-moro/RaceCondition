package main

import (
	"fmt"
	"sync/atomic"
	"time"

)

var num int32 = 0

func main() {
	

	for i := 0; i < 1000; i++ {
		go increment()
	}

     time.Sleep(time.Second * 5)

	fmt.Println(num)
}
func increment() {
	atomic.AddInt32(&num, 1)
}
