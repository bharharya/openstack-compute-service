package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Rate limiter settings
var limiter = rate.NewLimiter(5, 10) // 5 requests per second, burst size of 10

// RateLimitMiddleware limits the number of requests per client
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, slow down"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// SecureHeadersMiddleware adds security headers to all responses
func SecureHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	}
}
