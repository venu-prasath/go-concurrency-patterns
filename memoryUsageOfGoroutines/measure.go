package memoryusageofgoroutines

import (
	"fmt"
	"runtime"
	"sync"
)

func Measure() {
	memCounsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <- chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numOfGoroutines = 1e4
	wg.Add(numOfGoroutines)
	before := memCounsumed()
	for i:= numOfGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memCounsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numOfGoroutines/1000)
}	


