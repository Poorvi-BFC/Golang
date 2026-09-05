package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.RWMutex // TODO: Update mux to *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	// TODO: Lock for writing using sc.mux.Lock() and defer sc.mux.Unlock()
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	// TODO: Lock for reading using sc.mux.RLock() and defer sc.mux.RUnlock()
	sc.mux.RLock()
	defer sc.mux.RUnlock()
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
		mux:    &sync.RWMutex{},
	}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			sc.inc("emails")
			wg.Done()
		}()
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			_ = sc.val("emails")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Email count:", sc.val("emails"))
}
