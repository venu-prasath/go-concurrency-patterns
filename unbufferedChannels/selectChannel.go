package unbufferedChannels

/* The select statement in go allows a piece of code in any goroutine to listen to multiple channels.
 * The select statement blocks execution until it receives a value through any of the channels.
 * If any one of the channel emits a data, the select statement reads the value and unblocks the execution.
 * If multiple channels emit data, Go will randomly read from any of the channel.
 */

import "fmt"

func Select() {
	fmt.Println("\n\nHello from unbuffered-channel main()")
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		channelOne <- "data sent through channelOne"
	}()

	go func() {
		channelTwo <- "Data sent through channelTwo"
	}()

	select {
	case msgFromChannelOne := <- channelOne:
		fmt.Println(msgFromChannelOne)
	case msgFromChannelTwo := <- channelTwo:
		fmt.Println(msgFromChannelTwo)
	}

	fmt.Println("\nEnd of main() from unbuffered-channel")
}