package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readop struct {
	key  int
	resp chan int
}

type writeop struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0           // As before we’ll count how many operations we perform.
	reads := make(chan *readop) // The reads and writes channels will be used by other goroutines to issue read and write requests, respectively.
	writes := make(chan *writeop)

	go func() {
		var state = make(map[int]int) // now private to the stateful goroutine.
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ { // This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel.
		go func() {
			for {
				read := &readop{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeop{key: rand.Intn(5), val: rand.Intn(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	time.Sleep(time.Microsecond * 2)   // Let the goroutines work
	opsFinal := atomic.LoadInt64(&ops) // capture and report the ops count.
	fmt.Println("ops:", opsFinal)
}
