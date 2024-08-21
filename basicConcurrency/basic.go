package basicConcurrency

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