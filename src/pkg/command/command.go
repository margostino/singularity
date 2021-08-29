package command

import (
	"github.com/margostino/singularity/pkg/action"
	"github.com/margostino/singularity/pkg/config"
	"strings"
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

func (c *Command) WithArgs(args int) *Command {
	c.Args = args
	return c
}

func (c *Command) WithAction(action *action.Action) *Command {
	c.Action = action
	return c
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

func GetAction(command *config.CommandConfiguration) *action.Action {
	var commandAction *action.Action = nil
	if command.Args > 0 {
		function := action.InputActionStorage[command.Action]
		commandAction = action.NewInputAction(function)
	} else {
		function := action.ActionStorage[command.Action]
		commandAction = action.NewAction(function)
	}
	return commandAction
}

func InitializeCommands() *CommandMap {
	commands := make(map[string]*Command)
	configuration := config.GetCommandsConfiguration()

	for _, command := range configuration.CommandList {
		subcommands := strings.Split(command.Id, " ")
		var root *Command
		if commands[subcommands[0]] != nil {
			root = commands[subcommands[0]]
		} else {
			root = NewCommand(subcommands[0])
			commands[subcommands[0]] = root
		}

		lastIndex := len(subcommands[1:]) - 1
		current := root

		if lastIndex == -1 {
			action := GetAction(&command)
			root.WithArgs(command.Args).WithAction(action)
		}

		for i, subcommand := range subcommands[1:] {
			newCommand := createCommand(subcommand, i, lastIndex, &command)
			current.SubCommands[newCommand.Id] = newCommand
			current = newCommand
		}
	}
	return NewCommandMap(commands)
}

func createCommand(id string, currentIndex int, lastIndex int, command *config.CommandConfiguration) *Command {
	subcommand := NewCommand(id).WithArgs(command.Args)
	if currentIndex == lastIndex || lastIndex == -1 {
		action := GetAction(command)
		subcommand = subcommand.WithAction(action)
	}
	return subcommand
}

func Load() *CommandMap {
	return InitializeCommands()
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
