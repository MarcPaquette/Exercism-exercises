// Package gigasecond provide AddGigasecond to allow 10^9 seconds to be added to provided time
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 10^9 seconds to time t
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
