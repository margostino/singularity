package main

import (
	"github.com/margostino/singularity/pkg/config"
	"github.com/margostino/singularity/pkg/job"
	"github.com/margostino/singularity/pkg/preload"
	"github.com/margostino/singularity/pkg/runner"
	"github.com/margostino/singularity/pkg/shell"
)

func main() {
	config.LoadConfiguration()
	job.LoadJobs()
	shell.Welcome()
	preload.Preload()
	runner.Loop()
}
