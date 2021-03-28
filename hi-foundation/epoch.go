package main

import "fmt"
import "time"

// A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch.
// Epoch指的是一个特定的时间：1970-01-01 00:00:00 UTC
func main() { // Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds, respectively.
	now := time.Now()
	fmt.Println(now)

	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000 // Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.
	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	fmt.Println(time.Unix(secs, 0)) // You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	fmt.Println(time.Unix(0, nanos))
}
