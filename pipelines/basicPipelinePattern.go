package pipelines

/* The pipeline pattern is used when you want to process incoming data asynchronously,
 * but at the same time, pass the data between different processes through a channel
 * (a pipeline in this case). In the belew example, in the first goroutine, a slice of
 * integers are first extracted and pushed to a channel. A second goroutine reads from the
 * same channel to square the numbers and its output is pushed to a final channel.
 * The data is read from the final channel and printed as it arrives.
 */
import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int) 
	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}


func PipelinePattern() {
	fmt.Println("\n\nStart of PipelinePattern")
	nums := []int{2,3,4,5,6}

	//stage 1 - convert nums slice to a channel
	dataChannel := sliceToChannel(nums)

	//stage 2 - square every value in the dataChannel
	finalChannel := square(dataChannel)

	//stage 3 - read finalChannel
	for n := range finalChannel {
		fmt.Println(n)
	}
	fmt.Println("\nEnd of pipeline pattern")
}