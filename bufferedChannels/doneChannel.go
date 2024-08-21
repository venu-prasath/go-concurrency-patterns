package bufferedchannels

/* The done channel pattern is used to avoid a goroutine leak. It allows the parent goroutine
 * to cancel its children.
 * In the example below, the doWork function is called as a goroutine to run asyncrounously.
 * When the parent that started the doWork() closes the done channel, the goroutine stops execution
 * and returns. Not that, the done channel is passed as a readOnly channel, so that only the parent
 * can stop its execution.
 */

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing some work...")
		}
	}
}


func DoneChannelPattern() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(time.Second * 2)

	close(done)
}