package eveapi

import (
	"sync/atomic"
	"time"
)

type rateLimiter struct {
	throttle  chan time.Time
	rate      time.Duration
	ticker    *time.Ticker
	burstRate int
}

func newRateLimiter(requestsPerSecond int, burstRate int) *rateLimiter {
	c := &rateLimiter{
		throttle:  make(chan time.Time, burstRate),
		rate:      time.Second / (time.Duration)(requestsPerSecond),
		burstRate: burstRate,
	}
	c.startRatelimiter()
	return c
}

func (c *rateLimiter) startRatelimiter() {
	// Create the timed limiter
	c.ticker = time.NewTicker(c.rate)

	// Fill the buffer with the burst tokens
	for i := 0; i < c.burstRate; i++ {
		c.throttle <- time.Now()
	}

	// Start the rate limiter
	go c.tick()
}

func (c *rateLimiter) tick() {
	for t := range c.ticker.C {
		select {
		case c.throttle <- t:
		default:
		}
	}
}

func (c *rateLimiter) stop() {
	c.ticker.Stop()
}

func (c *rateLimiter) throttleRequest() {
	<-c.throttle
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
