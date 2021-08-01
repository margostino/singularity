package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"org.gene/singularity/pkg/commands"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/preload"
	"os"
	"strings"
	"time"
)

func main() {
	//go executeCronJob()
	commands.Welcome()
	config.LoadConfiguration()
	preload.Preload()
	commandTree := commands.Load()
	plan := Input()
	Validate(plan, commandTree)
	Loop(plan, commandTree)
}

func Loop(plan []string, commandTree *commands.CommandTree) {
	for {
		action := commandTree.Process(plan)
		action.Execute()
		plan = Input()
	}
}

func Validate(plan []string, commandTree *commands.CommandTree) {
	if !commandTree.IsValidPlan(plan) {
		fmt.Printf("command plan %q is not valid\n", plan)
		os.Exit(1)
	}
}

func Input() []string {
	commandLine := commands.Prompt()
	return strings.Fields(commandLine)
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
