package context

import (
	"fmt"
	"github.com/margostino/singularity/pkg/db"
	"time"
)

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

type WorldContext struct {
	State State
	Clock time.Time
}

type PlayerContext struct {
	Player *db.Player
}

var worldContext *WorldContext
var playerContext *PlayerContext

func NewContextBy(username string) {
	player := db.GetPlayerBy(username)
	playerContext = &PlayerContext{
		Player: player,
	}
}

func NewWorldContext() {
	worldContext = &WorldContext{
		State: Ready,
		Clock: time.Now(),
	}
}

func Print() {
	context := map[string]interface{}{
		"player": playerContext.Player.Username,
		"state":  worldContext.State,
		"clock":  worldContext.Clock,
	}
	fmt.Printf("%v\n", context)
}

func SetRunning() {
	// TODO: validate WorldContext == nil
	worldContext.State = Running
}

func UpdateWorldCycle() {
	// TODO: validate WorldContext == nil
	//worldContext.Clock = worldContext.Clock.AddDate(0, 0, 1)
	worldContext.Clock = worldContext.Clock.Add(time.Hour)
}

func GetUsername() string {
	if playerContext != nil && playerContext.Player != nil {
		return playerContext.Player.Username
	}
	return Guess
}

func GetState() State {
	if worldContext != nil {
		return worldContext.State
	}
	return Ready
}

func GetClock() string {
	if worldContext != nil {
		return worldContext.Clock.Format("2006-01-02 3 PM")
	}
	return time.Now().Format("2006-01-02 3 PM")
}

func Deactivate() {
	playerContext = nil
	worldContext.State = Ready
}

func IsPlayerSelected() bool {
	return playerContext != nil && playerContext.Player != nil
}

func Exit() {
	playerContext.Player = nil
}
