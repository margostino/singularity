package action

import (
	"fmt"
	"github.com/margostino/singularity/pkg/context"
	"os"
)

type Exit struct {
	Action
}

func ExecuteExit() {
	context.Deactivate()
	fmt.Println("bye!")
	os.Exit(0)
}
