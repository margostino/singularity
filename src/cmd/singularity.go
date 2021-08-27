package main

import (
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/job"
	"org.gene/singularity/pkg/preload"
	"org.gene/singularity/pkg/runner"
	"org.gene/singularity/pkg/shell"
)

func main() {
	config.LoadConfiguration()
	job.LoadJobs()
	shell.Welcome()
	preload.Preload()
	runner.Loop()
}
