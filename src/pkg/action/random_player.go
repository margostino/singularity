package action

import (
	"org.gene/singularity/pkg/context"
	"org.gene/singularity/pkg/db"
)

func ExecuteRandomPlayer() {
	player := db.PickPlayer()
	context.NewContextBy(player.Username)
}
