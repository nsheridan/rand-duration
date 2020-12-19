package randduration

import (
	"math/rand"
	"time"
)

var (
	defaultMin = time.Second * 10
	defaultMax = time.Minute * 5
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomDuration is a random duration generator.
type RandomDuration struct {
	Min, Max time.Duration
}

// Generate a random time.Duration between Min and Max, rounded to the nearest 100ms.
func (d *RandomDuration) Generate() time.Duration {
	min := d.Min
	max := d.Max
	if min <= 0 {
		min = defaultMin
	}
	if max <= 0 {
		max = defaultMax
	}
	if min >= max {
		// short circuit
		return max
	}
	mini := int64(min)
	maxi := int64(max)
	duration := time.Duration(rand.Int63n(maxi-mini+1) + mini)
	return duration.Round(100 * time.Millisecond)
}
