package memhog

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	// Seed the rand generator
	rand.Seed(time.Now().UnixNano())
}

// MemHog is an object that can hold data that is used to throttle RAM
type MemHog struct {
	Buffer [][]byte
}

// Return a []byte of size between range (min, max) in MiB
func makeBuffer(min, max int) []byte {
	i := min * 1048576
	j := max * 1048576
	return make([]byte, rand.Intn(j-i)+i)
}

func NewMemHog(bufferLen int) (*MemHog, error) {
	// Create a slice of byte slices with length: bufferSize
	b := make([][]byte, bufferLen)

	// Return a new MemHog
	return &MemHog{
		Buffer: b,
	}, nil
}

// Create a slice of buffers and randomly alloc its elements with data to
// increase & throttle RAM usage
func (t *MemHog) HogMemory() {
	var m runtime.MemStats
	bufferLen := len(t.Buffer)

	loops := 0
	for {
		loops += 1

		// Make a new buffer between 5-10 MiB, and set it into a
		// random index in the slice of buffers.
		// Any non-nil buffers that may be replaced by a new random buffer
		// essentially become disregarded, awaiting to be cleaned up by go's GC.
		i := rand.Intn(bufferLen)
		t.Buffer[i] = makeBuffer(5, 10)

		time.Sleep(1 * time.Second)

		bytes := 0
		for i := 0; i < bufferLen; i++ {
			if t.Buffer[i] != nil {
				// Compute the total bytes of the slice of buffers
				bytes += len(t.Buffer[i])
			}
		}

		// Print mem stats
		div := uint64(1048576) // 1048576 bytes = 1 MiB
		runtime.ReadMemStats(&m)
		fmt.Printf("heap: %d MiB\nbuffer_size: %d MiB\nheap_alloc: %d MiB\n"+
			"heap_idle: %d MiB\nheap_released: %d MiB\nloops: %d\n\n",
			m.HeapSys/div, bytes/int(div), m.HeapAlloc/div,
			m.HeapIdle/div, m.HeapReleased/div, loops)
	}
}
