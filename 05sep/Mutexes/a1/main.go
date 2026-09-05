package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	// TODO: Lock sc.mux and defer Unlock
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	// TODO: Lock sc.mux and defer Unlock
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	temp := sc.counts[key]
	time.Sleep(time.Microsecond)
	sc.counts[key] = temp + 1
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mux:    &sync.Mutex{},
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sc.inc("emails")
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Email count:", sc.val("emails"))
}
