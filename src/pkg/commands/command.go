package commands

type Command struct {
	Id          string
	SubCommands []*Command
	*Action
}

func NewCommand(id string) *Command {
	return &Command{
		Id:          id,
		SubCommands: make([]*Command, 0),
		Action:      nil,
	}
}

func (c Command) SubCommand(command *Command) *Command {
	c.SubCommands = append(c.SubCommands, command)
	return &c
}

func (c Command) WithAction(action *Action) *Command {
	c.Action = action
	return &c
}

func (c Command) Execute() {
	if c.Action != nil {
		c.Action.apply()
	}
}

func Load() *CommandTree {
	commands := make([]*Command, 0)
	show := createShowCommand()
	create := createCreateCommand()
	exit := createExitCommand()
	commands = append(commands, show, create, exit)
	return NewCommandTree(commands)
}

func createShowCommand() *Command {
	subCommandHelp := createShowHelpCommand()
	subCommandPlayers := createShowPlayersCommand()
	return NewCommand("show").SubCommand(subCommandHelp).SubCommand(subCommandPlayers)
}

func createExitCommand() *Command {
	exitAction := NewExitAction()
	return NewCommand("exit").WithAction(&exitAction.Action)
}

func createCreateCommand() *Command {
	subCommandPlayer := createCreatePlayerCommand()
	return NewCommand("create").SubCommand(subCommandPlayer)
}

func createCreatePlayerCommand() *Command {
	createPlayer := NewCreatePlayerAction()
	return NewCommand("players").WithAction(&createPlayer.Action)
}

func createShowHelpCommand() *Command {
	showHelp := NewShowHelpAction()
	return NewCommand("help").WithAction(&showHelp.Action)
}

func createShowPlayersCommand() *Command {
	showPlayers := NewShowPlayersAction()
	return NewCommand("players").WithAction(&showPlayers.Action)
}

func IsValidPlan(plan []string, commands []*Command) bool {
	for _, command := range commands {
		if command.Validate(plan) {
			return true
		}
	}
	return false
}

func (c Command) Validate(plan []string) bool {

	if (len(plan) > 0 && plan[0] != c.Id) ||
		(len(plan) == 1 && len(c.SubCommands) > 0) ||
		(len(plan) > 1 && len(c.SubCommands) == 0) {
		return false
	}

	if len(c.SubCommands) == 0 && plan[0] == c.Id && c.Action != nil {
		return true
	}

	for _, subCommand := range c.SubCommands {
		if subCommand.Validate(plan[1:]) {
			return true
		}
	}

	return false
}

func (c Command) GetAction(plan []string) *Action {

	if len(plan) == 0 {
		return nil
	}

	if len(plan) == 1 && plan[0] == c.Id && c.Action != nil {
		return c.Action
	}

	for _, subCommand := range c.SubCommands {
		return subCommand.GetAction(plan[1:])
	}

	return nil
}
