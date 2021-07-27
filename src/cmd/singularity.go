package main

import (
	"bufio"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/options"
	"org.gene/singularity/pkg/preload"
	"os"
	"time"
)

func main() {
	//go executeCronJob()
	config.LoadConfiguration()
	preload.Preload()
	options.PrintMainMenu()
	process()
	//SomeAPICallHandler() // handler which accepts requests and responds
}

func process() {
	var input = options.Input{Value: nil, Menu: options.NewMainMenu()}
	reader := bufio.NewReader(os.Stdin)
	for {
		option := options.ProcessInput(*reader)
		input.Value = &option
		input = options.ProcessOption(input)
	}
}

func myTask() {
	fmt.Println("This task will run periodically")
}
func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<-gocron.Start()
}

func SomeAPICallHandler() {
	time.Sleep(10000 * time.Millisecond)
}
