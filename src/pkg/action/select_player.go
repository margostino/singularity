package action

import (
	"org.gene/singularity/pkg/context"
)

func ExecuteSelectPlayer(args []string) {
	username := args[0]
	context.NewContext(username)
}
