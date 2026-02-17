package config

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
)

// SlidingWindowIncrementer wraps the Increment method.
type SlidingWindowIncrementer interface {
	// Increment increments the request counter for the current window and returns the counter values for the previous
	// window and the current one.
	// TTL is the time duration before the next window.
	Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (prevCount, currCount int64, err error)
}

// SlidingWindow implements a Sliding Window rate limiting algorithm.
//
// It does not require a distributed lock and uses a minimum amount of memory, however it will disallow all the requests
// in case when a client is flooding the service with requests.
// It's the client's responsibility to handle the disallowed request and wait before making a new request again.
type SlidingWindow struct {
	backend  SlidingWindowIncrementer
	clock    clock.Clock
	rate     time.Duration
	capacity int64
	epsilon  float64
}

// NewSlidingWindow creates a new instance of SlidingWindow.
// Capacity is the maximum amount of requests allowed per window.
// Rate is the window size.
// Epsilon is the max-allowed range of difference when comparing the current weighted number of requests with capacity.
func NewSlidingWindow(capacity int64, rate time.Duration, slidingWindowIncrementer SlidingWindowIncrementer, clock clock.Clock, epsilon float64) *SlidingWindow {
	return &SlidingWindow{backend: slidingWindowIncrementer, clock: clock, rate: rate, capacity: capacity, epsilon: epsilon}
}

// Limit returns the time duration to wait before the request can be processed.
// It returns ErrLimitExhausted if the request overflows the capacity.
func (s *SlidingWindow) Limit(ctx context.Context) (time.Duration, error) {
	now := s.clock.Now()
	currWindow := now.Truncate(s.rate)
	prevWindow := currWindow.Add(-s.rate)
	ttl := s.rate - now.Sub(currWindow)

	prev, curr, err := s.backend.Increment(ctx, prevWindow, currWindow, ttl+s.rate)
	if err != nil {
		return 0, err
	}

	// "prev" and "curr" are capped at "s.capacity + s.epsilon" using math.Ceil to round up any fractional values,
	// ensuring that in the worst case, "total" can be slightly greater than "s.capacity".
	prev = int64(math.Min(float64(prev), math.Ceil(float64(s.capacity)+s.epsilon)))
	curr = int64(math.Min(float64(curr), math.Ceil(float64(s.capacity)+s.epsilon)))

	total := float64(prev*int64(ttl))/float64(s.rate) + float64(curr)
	if total-float64(s.capacity) >= s.epsilon {
		var wait time.Duration
		if curr <= s.capacity-1 && prev > 0 {
			wait = ttl - time.Duration(float64(s.capacity-1-curr)/float64(prev)*float64(s.rate))
		} else {
			// If prev == 0.
			wait = ttl + time.Duration((1-float64(s.capacity-1)/float64(curr))*float64(s.rate))
		}

		return wait, errors.New("error limit exhausted")
		//ErrLimitExhausted
	}

	return 0, nil
}

func NewSlidingWindowInMemory() *SlidingWindowInMemory {
	return &SlidingWindowInMemory{}
}

type SlidingWindowInMemory struct {
	mu           sync.Mutex
	prevC, currC int64
	prevW, currW time.Time
}

// Increment increments the current window's counter and returns the number of requests in the previous window and the
// current one.
func (s *SlidingWindowInMemory) Increment(ctx context.Context, prev, curr time.Time, _ time.Duration) (int64, int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if curr != s.currW {
		if prev.Equal(s.currW) {
			s.prevW = s.currW
			s.prevC = s.currC
		} else {
			s.prevW = time.Time{}
			s.prevC = 0
		}

		s.currW = curr
		s.currC = 0
	}

	s.currC++

	return s.prevC, s.currC, ctx.Err()
}
