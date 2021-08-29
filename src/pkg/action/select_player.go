package action

import (
	"github.com/margostino/singularity/pkg/context"
)

func ExecuteSelectPlayer(args []string) {
	username := args[0]
	context.NewContextBy(username)
}
