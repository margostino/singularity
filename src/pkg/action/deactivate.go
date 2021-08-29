package action

import (
	"github.com/jasonlvhit/gocron"
	"org.gene/singularity/pkg/context"
)

func ExecuteDeactivate() {
	context.Deactivate()
	gocron.Clear()
}
