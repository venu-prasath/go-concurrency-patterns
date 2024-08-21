package main

import (
	"fmt"

	"github.com/venu-prasath/go-concurrency-patterns/basicConcurrency"
	"github.com/venu-prasath/go-concurrency-patterns/unbufferedChannels"
)




func main() {
	fmt.Println("Starting main...")

	//1. Basic concurrency Test
	basicConcurrency.BasicConcurrency()

	//2. Simple Unbuffered channel Test
	unbufferedChannels.SimpleChannel()

	//3. For Select pattern for multiple channels
	unbufferedChannels.Select()



	fmt.Println("\nEnd of main...")
}