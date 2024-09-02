package main

import (
	"fmt"
	"os"

	"github.com/venu-prasath/go-concurrency-patterns/basicConcurrency"
	bufferedchannels "github.com/venu-prasath/go-concurrency-patterns/bufferedChannels"
	memoryusageofgoroutines "github.com/venu-prasath/go-concurrency-patterns/memoryUsageOfGoroutines"
	ordonechannel "github.com/venu-prasath/go-concurrency-patterns/orDoneChannel"
	"github.com/venu-prasath/go-concurrency-patterns/pipelines"
	"github.com/venu-prasath/go-concurrency-patterns/unbufferedChannels"
	"github.com/venu-prasath/go-concurrency-patterns/util"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("No args provided")
		return
	}
	args := os.Args[1:]
	fmt.Println("Starting main... with args: ", args[0])

	switch args[0] {
	case "basicConcurrency":
		basicConcurrency.BasicConcurrency()
	case "simpleConcurrency":
		unbufferedChannels.SimpleChannel()
	case "selectPattern":
		unbufferedChannels.Select()
	case "forSelect":
		bufferedchannels.ForSelect()
	case "doneChannel":
		bufferedchannels.DoneChannelPattern()
	case "pipeline":
		pipelines.PipelinePattern()
	case "generator":
		util.Generator()
	case "intermediatePipe":
		pipelines.IntermediatePipeline()
	case "fanInFanOut":
		pipelines.FanOutPipeline()
	case "orDone":
		ordonechannel.OrDoneChannelPattern()
	case "memoryUsage":
		memoryusageofgoroutines.Measure()
	case "mutex":
		basicConcurrency.BasicMutex()
	case "rwMutex":
		basicConcurrency.RwMut()
	case "condition":
		basicConcurrency.ConditionCheck()
	case "cond_queue":
		basicConcurrency.QueueCondition()
	case "once":
		basicConcurrency.TestOnce()
	default:
		fmt.Println("Unknown argument")
		return
	}
	fmt.Println("\nEnd of main...")
}