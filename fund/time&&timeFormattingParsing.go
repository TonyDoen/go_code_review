package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	now := time.Now() // We’ll start by getting the current time.
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.
	p(then)
	p(then.Year()) // You can extract the various components of the time value as expected.
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())   // The Monday-Sunday Weekday is also available.
	p(then.Before(now)) // These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // The Sub methods returns a Duration representing the interval between two times.
	p(diff)               //
	p(diff.Hours())       // We can compute the length of the duration in various units.
	p(diff.Minutes())     //
	p(diff.Seconds())     //
	p(diff.Nanoseconds()) //
	p(then.Add(diff))     // You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	p(then.Add(-diff))

	// time formatting / parsing
	t := time.Now()
	p(t.Format(time.RFC3339))
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00") // Time parsing uses the same layout values as Format.
	p(t1)                                                          //
	p(t.Format("3:04PM"))                                          // Format and Parse use example-based layouts. Usually you’ll use a constant from time for these layouts, but you can also supply custom layouts. Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern with which to format/parse a given time/string. The example time must be exactly as shown: the year 2006, 15 for the hour, Monday for the day of the week, etc.
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()) // For purely numeric representations you can also use standard string formatting with the extracted components of the time value.
	ansic := "Mon Jan _2 15:04:05 2006"                                                                               // Parse will return an error on malformed input explaining the parsing problem.
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
