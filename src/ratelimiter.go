package src

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	cooldowns map[string]time.Time
	duration  time.Duration
}

func NewRateLimiter(duration time.Duration) *RateLimiter {
	return &RateLimiter{
		cooldowns: make(map[string]time.Time),
		duration:  duration,
	}
}

// CheckAndUpdate returns true if the key can proceed (not rate limited)
// and updates the cooldown time if proceeding
func (rl *RateLimiter) CheckAndUpdate(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	lastUsed, exists := rl.cooldowns[key]

	if exists && now.Sub(lastUsed) < rl.duration {
		return false // Rate limited
	}

	rl.cooldowns[key] = now
	return true // Can proceed
}

// GetRemainingCooldown returns the remaining cooldown duration for a key
func (rl *RateLimiter) GetRemainingCooldown(key string) time.Duration {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	lastUsed, exists := rl.cooldowns[key]
	if !exists {
		return 0
	}

	elapsed := time.Since(lastUsed)
	if elapsed >= rl.duration {
		return 0
	}

	return rl.duration - elapsed
}