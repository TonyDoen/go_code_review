package main

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the future, or repeatedly at some interval. Go’s built-in timer and ticker features make both of these tasks easy.
func main() {
	// timer
	timer1 := time.NewTimer(time.Second * 1) // Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
	<-timer1.C                               // The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	fmt.Println("timer 1 expired.")

	// One reason a timer may be useful is that you can cancel the timer before it expires.
	timer2 := time.NewTimer(time.Second * 4)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired.")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	//  tickers are for when you want to do something repeatedly at regular intervals.
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
