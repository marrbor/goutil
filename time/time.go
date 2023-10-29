package time

import "time"

// JST returns pointer of time.Location instance that points Japan Standard Time.
func JST() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

// WaitSec waited given second(s).
func WaitSec(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

// WaitMsec waited given millisecond(s).
func WaitMsec(msec int) {
	time.Sleep(time.Duration(msec) * time.Millisecond)
}
