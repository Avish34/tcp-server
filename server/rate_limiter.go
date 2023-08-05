package server

import (
	"log"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	MaxTokens 			int64
	CurTokens 			int64
	Rate      			int64
	LastRefillTimeStamp time.Time
	Mutex               *sync.Mutex
}

func NewRateLimiter(rate, tokens int64) TokenBucket {
	return TokenBucket{
		MaxTokens: tokens,
		Rate: rate,
		LastRefillTimeStamp: time.Now(),
		Mutex: &sync.Mutex{},
	}
}

func (tb *TokenBucket) refillBucket() {
	log.Println("Refilling the bucket")
	now := time.Now()
	last := time.Since(tb.LastRefillTimeStamp)
	tokensToBeAdded := (last.Milliseconds() * tb.Rate) / 1000
	log.Printf("Adding %d tokens to bucket at %v", tokensToBeAdded, now)
	tb.CurTokens = int64(math.Min(float64(tb.CurTokens + tokensToBeAdded), float64(tb.MaxTokens)))
	tb.LastRefillTimeStamp = now
}

func (tb *TokenBucket) isRequestAllowed() bool {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()
	tb.refillBucket()
	if tb.CurTokens > 0 {
		tb.CurTokens --
		return true
	}
	return false
}