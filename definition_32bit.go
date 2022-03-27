//go:build (386 || arm || mipsle || mips || wasm)

package backoff

import (
	"math"
	"sync/atomic"
    "time"
)

const maxDuration = float64(math.MaxInt32 - 512)

// Backoff is a time.Duration counter, starting at Min. After every call to
// the Duration method the current timing is multiplied by Factor, but it
// never exceeds Max.
//
// Backoff is not generally concurrent-safe, but the ForAttempt method can
// be used concurrently.
type Backoff struct {
	attempt uint32
	// Factor is the multiplying factor for each increment step
	Factor float64
	// Jitter eases contention by randomizing backoff steps
	Jitter bool
	// Min and Max are the minimum and maximum values of the counter
	Min, Max time.Duration
}

// Duration returns the duration for the current attempt before incrementing
// the attempt counter. See ForAttempt.
func (b *Backoff) Duration() time.Duration {
	d := b.ForAttempt(float64(atomic.AddUint32(&b.attempt, 1) - 1))
	return d
}

// Reset restarts the current attempt counter at zero.
func (b *Backoff) Reset() {
	atomic.StoreUint32(&b.attempt, 0)
}

// Attempt returns the current attempt counter value.
func (b *Backoff) Attempt() float64 {
	return float64(atomic.LoadUint32(&b.attempt))
}
