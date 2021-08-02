package commands

import (
	"fmt"
	"org.gene/singularity/pkg/db"
)

func ExecuteShowPlayers() {
	fmt.Printf("%v\n", db.Players)
}
