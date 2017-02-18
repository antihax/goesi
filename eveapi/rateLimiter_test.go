package eveapi

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	r := newRateLimiter(1, 20)
	count := 0
	start := time.Now()
	for {
		count++
		r.throttleRequest()
		if count == 20 {
			duration:=time.Now().Sub(start).Seconds()
			if (int)(duration) != 0 {
				t.Errorf("Burst failed") // This should be immediate
			}
		}
		if count == 22 {
			duration:=time.Now().Sub(start).Seconds()
			if (int)(duration) != 2 { // after two seconds
				t.Errorf("Rate limiter failed to properly limit %d", duration)
			}
			time.Sleep(time.Second * 2)
		}

		if count == 25 {
			duration:=time.Now().Sub(start).Seconds()
			if (int)(duration) != 5 { // after five seconds
				t.Errorf("Failed to recover burst tokens %d", duration)
			}
			break
		}
	}
}
