package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
func main() {
	var ops uint64 = 0        // We’ll use an unsigned integer to represent our (always-positive) counter.
	for i := 0; i < 50; i++ { // To simulate concurrent updates, we’ll start 50 goroutines that each increment the counter about once a millisecond.
		go func() {
			for {
				atomic.AddUint64(&ops, 1) // To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
				runtime.Gosched()         // Allow other goroutines to proceed. // runtime.Gosched()用于让出CPU时间片。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。
			}
		}()
	}
	time.Sleep(time.Second)             // Wait a second to allow some ops to accumulate.
	opsFinal := atomic.LoadUint64(&ops) // In order to safely use the counter while it’s still being updated by other goroutines, we extract a copy of the current value into opsFinal via LoadUint64.
	fmt.Println("ops:", opsFinal)
}
