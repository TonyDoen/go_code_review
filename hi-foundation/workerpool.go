package main

import (
	"fmt"
	"time"
)

// Here’s the worker, of which we’ll run several concurrent instances. These workers will receive work on the jobs channel and send the corresponding results on results. We’ll sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second * 1)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w < 4; w++ { // This starts up 3 workers, initially blocked because there are no jobs yet.
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ { // Here we send 9 jobs and then close that channel to indicate that’s all the work we have.
		jobs <- j
	}
	close(jobs)

	for a := 1; a < 10; a++ { // Finally we collect all the results of the work.
		<-results
	}
}
