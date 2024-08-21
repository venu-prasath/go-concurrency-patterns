package main

import (
	"fmt"

	ordonechannel "github.com/venu-prasath/go-concurrency-patterns/orDoneChannel"
)




func main() {
	fmt.Println("Starting main...")

	//1. Basic concurrency Test
	//basicConcurrency.BasicConcurrency()

	//2. Simple Unbuffered channel Test
	//unbufferedChannels.SimpleChannel()

	//3. Select pattern for multiple channels
	//unbufferedChannels.Select()

	//4. For Select pattern for multiple channels
	//bufferedchannels.ForSelect()

	//5. Done Channel with For Select
	//bufferedchannels.DoneChannelPattern()

	//6. Basic Pipeline
	//pipelines.PipelinePattern()

	//7. Generator
	//util.Generator()

	//8. Reading fixed size values from a generator
	//pipelines.IntermediatePipeline()

	//9. Fan out Fan in pipeline
	//pipelines.FanOutPipeline()

	//10. OrDoneChannel Pattern
	ordonechannel.OrDoneChannelPattern()


	fmt.Println("\nEnd of main...")
}