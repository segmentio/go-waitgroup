package waitgroup_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/segmentio/go-waitgroup"
)

func TestWaitTicksWhenDoneIsCalled(t *testing.T) {
	var wg waitgroup.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			wg.Done()
		}()
	}

	select {
	case <-wg.Wait():
	case <-time.After(1 * time.Millisecond):
		t.Errorf("wg.Wait() channel should have ticked")
	}
}

func TestWaitDoesNotTickWhenDoneIsnotCalled(t *testing.T) {
	var wg waitgroup.WaitGroup
	wg.Add(5)

	select {
	case <-wg.Wait():
		t.Errorf("wg.Wait() channel should not have ticked")
	case <-time.After(1 * time.Millisecond):
	}
}

func ExampleWait() {
	var wg waitgroup.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println(i)
		}(i)
	}

	select {
	// Try reducing this to trigger a timeout.
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	case <-wg.Wait():
		fmt.Println("done")
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// done
}
