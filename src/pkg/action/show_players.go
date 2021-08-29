package action

import (
	"fmt"
	"github.com/margostino/singularity/pkg/db"
)

func ExecuteShowPlayers() {
	fmt.Printf("%v\n", db.Players)
}
