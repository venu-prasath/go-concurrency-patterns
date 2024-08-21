package unbufferedChannels

/* Channels are a typed, thead-safe queue. Channels allow different goroutines to
 * communicate with each other. The <- operator is called the channel operator.
 * Data flows in the direction of the arrow. This operation will block until another
 * goroutine is ready to receive the value
 */

import "fmt"


func SimpleChannel() {
	fmt.Println("\n\nStart from SimpleChannel() from unbufferedChannel")
	myChannel := make(chan string)

	go func() {
		myChannel <- "data sent through channel from a go-routine"
	}()

	msg := <- myChannel

	fmt.Println("Msg received from channel in main: ", msg)

	fmt.Println("\nEnd of SimpleChannel() from unbufferedChannel")
}