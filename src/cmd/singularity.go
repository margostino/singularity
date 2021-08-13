package main

import (
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/preload"
	"org.gene/singularity/pkg/runner"
	"org.gene/singularity/pkg/shell"
)

func main() {
	config.LoadConfiguration()
	shell.Welcome()
	preload.Preload()
	runner.Loop()
}
