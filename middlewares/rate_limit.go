package middlewares

import (
    "net/http"
    "sync"
    "golang.org/x/time/rate"
    "github.com/gin-gonic/gin"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()

    limiter, exists := visitors[ip]
    if !exists {
        limiter = rate.NewLimiter(5, 10)
        visitors[ip] = limiter
    }

    return limiter
}

func RateLimit() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        limiter := getVisitor(ip)

        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
                "error": "Too many requests. Please slow down.",
            })
            return
        }

        c.Next()
    }
}
