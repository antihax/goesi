package eveapi

import (
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

type rateLimiter struct {
	lim *rate.Limiter
}

func newRateLimiter(requestsPerSecond int, burstRate int) *rateLimiter {
	lim := rate.NewLimiter(rate.Limit(requestsPerSecond), burstRate)
	return &rateLimiter{lim}
}

func (c *rateLimiter) throttleRequest() {
	r := c.lim.Reserve()
	time.Sleep(r.Delay())
}

type concurrencyLimiter struct {
	concurrencyLimiter chan bool
	openRequests       uint64
}

func newConcurrencyLimiter(limit int) *concurrencyLimiter {
	c := &concurrencyLimiter{make(chan bool, limit), 0}
	return c
}

func (c *concurrencyLimiter) startRequest() {
	c.concurrencyLimiter <- true
	atomic.AddUint64(&c.openRequests, 1)
}

func (c *concurrencyLimiter) endRequest() {
	<-c.concurrencyLimiter
	atomic.AddUint64(&c.openRequests, ^uint64(0))
}

func (c *concurrencyLimiter) getOpenRequests() uint64 {
	return atomic.LoadUint64(&c.openRequests)
}

var connectionLimit *concurrencyLimiter

// CCP's documentation states rate limits are tracked by IP address.
// The throttles provide a burstable rate limit to each component of the API.
var authedThrottle *rateLimiter
var anonThrottle *rateLimiter
var xmlThrottle *rateLimiter

func init() {
	// Rate limits
	authedThrottle = newRateLimiter(20, 100)
	anonThrottle = newRateLimiter(150, 400)
	xmlThrottle = newRateLimiter(30, 30)

	// Prevent going over 20 concurrent requests
	connectionLimit = newConcurrencyLimiter(20)
}
