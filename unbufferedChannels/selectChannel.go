package unbufferedChannels

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