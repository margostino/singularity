package action

import (
	"github.com/margostino/singularity/pkg/context"
	"github.com/margostino/singularity/pkg/db"
)

func ExecuteRandomPlayer() {
	player := db.PickPlayer()
	context.NewContextBy(player.Username)
}
