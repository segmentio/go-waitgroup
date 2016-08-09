// Package waitgroup provides a WaitGroup with a channel based Wait() method.
package waitgroup

import "sync"

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to monitor when goroutines have finished.
type WaitGroup struct {
	wg sync.WaitGroup
}

// Add adds delta, which may be negative, to the WaitGroup counter.
// If the counter becomes zero, all goroutines blocked on Wait are released.
// If the counter goes negative, Add panics.
// See https://golang.org/pkg/sync/#WaitGroup.Add.
func (wg *WaitGroup) Add(delta int) {
	wg.wg.Add(delta)
}

// Done decrements the WaitGroup counter. See https://golang.org/pkg/sync/#WaitGroup.Done.
func (wg *WaitGroup) Done() {
	wg.wg.Done()
}

// Wait returns a channel that emits a tick when the WaitGroup counter is zero.
// Equivalent to https://golang.org/pkg/sync/#WaitGroup.Wait but uses a channel
// instead of blocking.
func (wg *WaitGroup) Wait() <-chan struct{} {
	out := make(chan struct{})
	go func() {
		wg.wg.Wait()
		out <- struct{}{}
		close(out)
	}()
	return out
}
