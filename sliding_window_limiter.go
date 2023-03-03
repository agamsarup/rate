package rate

import (
	"sync"
	"time"
)

type SlidingWindowRateLimiter struct {
	mu                  sync.Mutex
	maxAllowedRequests  int
	windowSize          time.Duration
	curWindowStartTime  time.Time
	memoryStore         SlidingWindowMemoryStore
	curRequestTimestamp time.Time
}

func NewSlidingWindowRateLimiter(maxAllowedRequests int, windowSize time.Duration, memoryStore SlidingWindowMemoryStore) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		maxAllowedRequests: maxAllowedRequests,
		windowSize:         windowSize,
		curWindowStartTime: time.Now(),
		memoryStore:        memoryStore,
	}
}

func (r *SlidingWindowRateLimiter) Allow(id string) bool {
	r.curRequestTimestamp = time.Now()

	r.mu.Lock()
	defer r.mu.Unlock()

	r.updateWindowsIfApplicable()

	r.memoryStore.Increment(id, 1)
	prevWindowContibutionFactor := float64(1) - float64(r.curRequestTimestamp.Sub(r.curWindowStartTime))/float64(r.windowSize)
	prevWindowRequests := float64(r.memoryStore.PreviousWindowCount(id))
	curWindowRequests := float64(r.memoryStore.CurrentWindowCount(id))

	effectiveRequestsByWindow := prevWindowContibutionFactor*prevWindowRequests + curWindowRequests

	return effectiveRequestsByWindow <= float64(r.maxAllowedRequests)
}

func (r *SlidingWindowRateLimiter) updateWindowsIfApplicable() {
	timeDiff := r.curRequestTimestamp.Sub(r.curWindowStartTime)
	if timeDiff > 2*r.windowSize {
		r.curWindowStartTime = r.curRequestTimestamp
		r.memoryStore.Reset()
	} else if timeDiff > r.windowSize {
		r.curWindowStartTime = r.curWindowStartTime.Add(r.windowSize)
		r.memoryStore.MoveToNextWindow()
	}
}
