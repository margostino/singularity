package context

import "org.gene/singularity/pkg/db"

type State int

const (
	Running State = iota
	Ready
	Stopped
	Paused
)

func (s State) String() string {
	return [...]string{"Running", "Ready", "Stopped", "Paused"}[s]
}

type GameContext struct {
	Player *db.Player
	State  State
}

var WorldContext *GameContext

func NewContext(username string) {
	player := db.GetPlayerBy(username)
	WorldContext = &GameContext{
		Player: player,
		State:  Ready,
	}
}

func SetRunning() {
	// TODO: validate WorldContext == nil
	WorldContext.State = Running
}

func GetUsername() string {
	if WorldContext != nil && WorldContext.Player != nil {
		return WorldContext.Player.Username
	}
	return "guess"
}

func GetState() State {
	if WorldContext != nil {
		return WorldContext.State
	}
	return Ready
}

func Deactivate() {
	WorldContext = nil
}
