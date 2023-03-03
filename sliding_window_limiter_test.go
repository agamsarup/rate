package rate

import (
	"testing"
	"time"
)

func TestSlidingWindowLimiter(t *testing.T) {
	limiter := NewSlidingWindowRateLimiter(10, time.Second, NewSlidingWindowInMemoryStore())
	ch := sendRequests()

	for req := range ch {
		allow := limiter.Allow("1")
		//fmt.Printf("req=%d allow=%t\n", req, allow)
		if (req <= 10 && !allow) || (req > 10 && allow) {
			t.FailNow()
		}
		if req == 20 {
			break
		}
	}
	time.Sleep(1500 * time.Millisecond)

	for req := range ch {
		allow := limiter.Allow("1")
		//fmt.Printf("req=%d allow=%t\n", req, allow)
		if (req <= 25 && !allow) || (req > 25 && allow) {
			t.FailNow()
		}
	}
}

func sendRequests() chan int {
	reqChan := make(chan int)
	go func() {
		for i := 1; i <= 40; i++ {
			reqChan <- i
		}
		close(reqChan)
	}()
	return reqChan
}
