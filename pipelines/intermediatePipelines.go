package pipelines

/* This pattern is an intermediate implementation of the basic pipeline, where I am simulating a slow running
 * pipeline by finding if a large number is prime or not in this case. There are many real world use cases
 * such as ML, Image processing, video processing pipelines which needs to happen asychrounously.
 * The code snippet shows a naive way of building a single pipeline that waits for the completion of a slow running
 * task to output the result.
 */

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/venu-prasath/go-concurrency-patterns/util"
)



func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <- done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}


// Simulating a slow pipeline stage
func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i>1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				//fmt.Println("Recieved large number: ", randomInt)
				if isPrime(randomInt) {
					//fmt.Println("Found Prime!!!")
					primes <- randomInt 
				}
			}
		}

	}()
	return primes
}

func IntermediatePipeline() {
	start := time.Now()
	done := make(chan int) 
	defer close(done)

	randomNumFetcher := func() int { return rand.Intn(5000000) }

	randIntStream := util.RepeatFunc(done, randomNumFetcher)

	primeStream := primeFinder(done, randIntStream)

	for rando := range take(done, primeStream, 10) {
		fmt.Println(rando)
	}
	fmt.Println(time.Since(start))
}