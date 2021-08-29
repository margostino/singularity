package action

import (
	"github.com/jasonlvhit/gocron"
	"github.com/margostino/singularity/pkg/context"
)

func ExecuteDeactivate() {
	context.Deactivate()
	gocron.Clear()
}
