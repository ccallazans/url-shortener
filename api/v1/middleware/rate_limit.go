package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const MAX_REQUESTS_PER_MINUTE = 10

type IpLimiter struct {
	ips map[string]*rateLimiter
	mu  sync.Mutex
}

type rateLimiter struct {
	lastSeen time.Time
	hits     int
}

func NewIpLimiter() *IpLimiter {
	return &IpLimiter{ips: make(map[string]*rateLimiter)}
}

func (il *IpLimiter) Increment(ip string) {
	il.mu.Lock()
	defer il.mu.Unlock()

	rl, ok := il.ips[ip]
	if !ok {
		rl = &rateLimiter{}
		il.ips[ip] = rl
	}

	rl.hits++
	rl.lastSeen = time.Now()
}

func (il *IpLimiter) CleanUp() {
	il.mu.Lock()
	defer il.mu.Unlock()

	for ip, rl := range il.ips {
		if time.Since(rl.lastSeen) > time.Hour {
			delete(il.ips, ip)
		}
	}
}

func LimitIpRequests() gin.HandlerFunc {
	il := NewIpLimiter()

	go func() {
		for range time.Tick(time.Minute) {
			il.CleanUp()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		il.Increment(ip)

		rl, ok := il.ips[ip]
		if ok && rl.hits > MAX_REQUESTS_PER_MINUTE {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests from your IP address. Maximum of 10 creations per minute. Try it again later.",
			})
			return
		}

		c.Next()
	}
}
