package util

/* A generator is a code to generate a stream of data and push it to a channel.
 * This is used to simulate realtime data received by a backend application from
 * real time systems. It could be time-series data read from a financial system.
 * It could be real time data read from IOT devices. It could be real time data
 * received from chat applications or MMORPG games, etc.
 */

import (
	"fmt"
	"math/rand"
)

func RepeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <- done:
				return
			case stream <- fn():
			}
		}

	}()
	return stream
}

func Generator() {
	done := make(chan int) 
	defer close(done)

	randomNumFetcher := func() int { return rand.Intn(500000000) }

	for rando := range RepeatFunc(done, randomNumFetcher) {
		fmt.Println(rando)
	}
}