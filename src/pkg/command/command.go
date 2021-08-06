package command

import (
	"org.gene/singularity/pkg/action"
)

type Command struct {
	Id          string
	Args        int
	SubCommands map[string]*Command
	*action.Action
}

func NewCommand(id string) *Command {
	return &Command{
		Id:          id,
		Args:        0,
		SubCommands: make(map[string]*Command),
		Action:      nil,
	}
}

func (c Command) SubCommand(command *Command) *Command {
	c.SubCommands[command.Id] = command
	return &c
}

func (c Command) WithArgs(args int) *Command {
	c.Args = args
	return &c
}

func (c Command) WithAction(action *action.Action) *Command {
	c.Action = action
	return &c
}

func (c Command) Execute() {
	if c.Action != nil {
		c.Action.Function()
	}
}

func (c Command) ExecuteWith(args []string) {
	if c.Action != nil {
		c.Action.InputFunction(args)
	}
}

func Load() *CommandMap {
	commands := make(map[string]*Command, 0)
	show := createShowCommand()
	create := createCreateCommand()
	exit := createCommand("exit", action.NewAction(action.ExecuteExit))
	start := createCommand("start", action.NewAction(action.ExecuteStart))
	selectCo := createSelectCommand()
	commands = map[string]*Command{
		"show":   show,
		"create": create,
		"exit":   exit,
		"start":  start,
		"select": selectCo,
	}
	return NewCommandMap(commands)
}

func Load2() *CommandTree {
	commands := make([]*Command, 0)
	show := createShowCommand()
	create := createCreateCommand()
	exit := createCommand("exit", action.NewAction(action.ExecuteExit))
	start := createCommand("start", action.NewAction(action.ExecuteStart))
	selectCo := createSelectCommand()
	commands = append(commands, show, create, exit, start, selectCo)
	return NewCommandTree(commands)
}

func createSelectCommand() *Command {
	player := createCommand("player", action.NewInputAction(action.ExecuteSelectPlayer)).WithArgs(1)
	return NewCommand("select").
		SubCommand(player)
}

func createShowCommand() *Command {
	help := createCommand("help", action.NewAction(action.ExecuteShowHelp))
	players := createCommand("players", action.NewAction(action.ExecuteShowPlayers))
	stats := createCommand("stats", action.NewAction(action.ExecuteShowStats))
	return NewCommand("show").
		SubCommand(help).
		SubCommand(players).
		SubCommand(stats)
}

func createCreateCommand() *Command {
	player := createCommand("player", action.NewAction(action.ExecuteCreatePlayer))
	return NewCommand("create").
		SubCommand(player)
}

func createCommand(id string, action *action.Action) *Command {
	return NewCommand(id).WithAction(action)
}

func (c Command) isLastCommand(plan []string) bool {
	return len(plan)-c.Args == 1 && plan[0] == c.Id && c.Action != nil && len(c.SubCommands) == 0
}

func (c Command) Validate(plan []string) bool {
	command := c.GetLastCommand(plan)
	return command != nil
}

func (c Command) GetLastCommand(plan []string) *Command {

	if len(plan) == 0 || plan[0] != c.Id {
		return nil
	}

	if c.isLastCommand(plan) {
		return &c
	}

	if len(plan) > 1 {
		command, ok := c.SubCommands[plan[1]]
		if ok {
			return command.GetLastCommand(plan[1:])
		}
	}

	return nil
}
