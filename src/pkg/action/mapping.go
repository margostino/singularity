package action

var ActionStorage = map[string]func(){
	"ExecuteExit":         ExecuteExit,
	"ExecuteStart":        ExecuteStart,
	"ExecuteDeactivate":   ExecuteDeactivate,
	"ExecuteShowPlayers":  ExecuteShowPlayers,
	"ExecuteShowStats":    ExecuteShowStats,
	"ExecuteHelp":         ExecuteHelp,
	"ExecuteCreatePlayer": ExecuteCreatePlayer,
}

var InputActionStorage = map[string]func([]string){
	"ExecuteSelectPlayer": ExecuteSelectPlayer,
}
