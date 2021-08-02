package commands

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"math/rand"
	"org.gene/singularity/pkg/db"
	"time"
)

func ExecuteStart() {
	fmt.Println("start!")
	go executeCronJob()
}

func myTask() {
	//fmt.Println("This task will run periodically")
	value := rand.Intn(100)
	country := db.PickCountry()
	country.WarmingMetrics[0].Value = value
}

func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<-gocron.Start()
}

func SomeAPICallHandler() {
	time.Sleep(10000 * time.Millisecond)
}
