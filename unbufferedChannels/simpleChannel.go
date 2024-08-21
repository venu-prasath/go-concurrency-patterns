package unbufferedChannels

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