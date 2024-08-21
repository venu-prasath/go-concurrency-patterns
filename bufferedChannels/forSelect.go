package bufferedchannels

/* The for select pattern of concurrency, is to continuously read data from a channel.
 * This pattern is useful when you want a long running service to continuously read from
 * a channel throughtout the lifetime of the application.
 * A done channel is used with this pattern to control the execution of a goroutine.
 */

import "fmt"

func ForSelect() {
	fmt.Println("\n\nStart of ForSelect")
	charChannel := make(chan string, 3)

	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:

		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	fmt.Println("\nEnd of ForSelect")
}