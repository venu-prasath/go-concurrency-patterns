package ordonechannel

/* This is an example pattern to showcase the orDone pattern. This cleans up the code
 * by removing repeated for select done logic in one place.
 */

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup


func OrDoneChannelPattern() {
	done := make(chan interface{})
	defer close(done)

	channelOne := make(chan interface{}, 100)
	channelTwo := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-done:
				return
			case channelOne <- "Data for channel 1":
			}

		}

	}()

	go func() {
		for {
			select {
			case <- done:
				return
			case channelTwo <- "Data for channel 2":
			}
		}
	}()

	wg.Add(1)
	go consumeChannelOne(done, channelOne)
	wg.Add(1)
	go consumeChannelTwo(done, channelTwo)

	wg.Wait()

}

func consumeChannelOne(done <-chan interface{}, chOne <-chan interface{}) {
	defer wg.Done()
	for val := range orDone(done, chOne) {
		fmt.Println(val)
		//Other important complex logic related to consumer 1
	}
}

func consumeChannelTwo(done <-chan interface{}, chTwo <-chan interface{}) {
	defer wg.Done()
	for val := range orDone(done, chTwo) {
		fmt.Println(val)
		//Other important complex logic related to consumer 2
	}	
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return relayStream
}