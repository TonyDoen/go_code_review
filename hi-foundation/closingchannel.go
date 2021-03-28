package main

import "fmt"

func main() {
	//closing channels
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() { // worker
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j < 4; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // This sends 3 jobs to the worker over the jobs channel, then closes it.
	fmt.Println("sent all jobs")

	<-done // We await the worker using the synchronization approach we saw earlier.

	//range over channels
	queue := make(chan string, 3) // This example showed that it’s possible to close a non-empty channel but still have the remaining values be received.
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)
	for elem := range queue { //  We can also use "for and range" syntax to iterate over values received from a channel.
		fmt.Println("element:", elem) // Because we closed the channel above, the iteration terminates after receiving the 3 elements. If we didn’t close it we’d block on a 4th receive in the loop.
	}
}
