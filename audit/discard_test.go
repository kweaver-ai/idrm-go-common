package audit

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Logger instance produced by Discard() always compare as equal.
func TestDiscard(t *testing.T) {
	logger := Discard()

	const n = 1 << 6

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()

			for j := 0; j < n; j++ {
				assert.Equal(t, logger, Discard(), "process %d: %d", i, j)
			}
		}(i)
	}
	wg.Wait()
}
