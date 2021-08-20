package context

import "org.gene/singularity/pkg/db"

type State int

const (
	Running State = iota
	Ready
	Stopped
	Paused
	Guess = "guess"
)

func (s State) String() string {
	return [...]string{"Running", "Ready", "Stopped", "Paused"}[s]
}

type GameContext struct {
	Player *db.Player
	State  State
}

var worldContext *GameContext

func NewContext(username string) {
	player := db.GetPlayerBy(username)
	worldContext = &GameContext{
		Player: player,
		State:  Ready,
	}
}

func SetRunning() {
	// TODO: validate WorldContext == nil
	worldContext.State = Running
}

func GetUsername() string {
	if worldContext != nil && worldContext.Player != nil {
		return worldContext.Player.Username
	}
	return Guess
}

func GetState() State {
	if worldContext != nil {
		return worldContext.State
	}
	return Ready
}

func Deactivate() {
	worldContext = nil
}

func IsPlayerSelected() bool {
	return worldContext.Player != nil
}

func Exit() {
	worldContext.Player = nil
}
