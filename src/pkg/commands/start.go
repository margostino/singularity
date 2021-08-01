package commands

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func ExecuteStart() {
	fmt.Println("start!")
	go executeCronJob()
}

func myTask() {
	//fmt.Println("This task will run periodically")
}
func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<-gocron.Start()
}

func SomeAPICallHandler() {
	time.Sleep(10000 * time.Millisecond)
}