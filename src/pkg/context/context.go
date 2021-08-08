package context

import "org.gene/singularity/pkg/db"

type GameContext struct {
	Player *db.Player
}

var WorldContext *GameContext

func NewContext(username string) {
	player := db.GetPlayerBy(username)
	WorldContext = &GameContext{Player: player}
}

func GetUsername() string {
	if WorldContext != nil && WorldContext.Player != nil {
		return WorldContext.Player.Username
	}
	return ""
}

func Deactivate() {
	WorldContext = nil
}
