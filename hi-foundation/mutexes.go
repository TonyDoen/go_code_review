package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//  For more complex state we can use a mutex to safely access data across multiple goroutines.
func main() {
	var state = make(map[int]int) // For our example the state will be a map.
	var mutex = &sync.Mutex{}     // This mutex will synchronize access to state.
	var ops int64 = 0

	// read
	for r := 0; r < 100; r++ { // Here we start 100 goroutines to execute repeated reads against the state.
		go func() {
			total := 0
			for {
				key := rand.Intn(5)      // For each read we pick a key to access,
				mutex.Lock()             // Lock() the mutex to ensure exclusive access to the state,
				total += state[key]      //
				mutex.Unlock()           // Unlock() the mutex, and increment the ops count.
				atomic.AddInt64(&ops, 1) //
				runtime.Gosched()        // In order to ensure that this goroutine doesn’t starve the scheduler, we explicitly yield after each operation with runtime.Gosched().
			}
		}()
	}
	time.Sleep(time.Microsecond * 2)
	opsfinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsfinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

	// write
	for w := 0; w < 10; w++ { // start 10 goroutines to simulate writes, using the same pattern we did for reads.
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Microsecond * 1)
	opsfinal = atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsfinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
