package basicConcurrency

/* The main idea behing goroutines is to run async task in a lightweight
* thread in the background.
* Go was designed to be concurrent. It excels at performing many tasks simultaneously
* safely using a simple syntax.
* Goroutines are light-weight. It's memory footprint is lighter that OS-level threads.
 */

import (
	"fmt"
	"time"
)

func simpleFunc(num int) {
	fmt.Println("The num received from argument: ", num)
}

func BasicConcurrency() {
	fmt.Println("\n\nStart of basic concurrency!")
	go simpleFunc(1)
	go simpleFunc(2)
	go simpleFunc(3)

	time.Sleep(time.Second * 3)
	fmt.Println("End of basic concurrency!")
}