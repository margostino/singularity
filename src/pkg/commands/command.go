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
	exit := createCommand("exit", NewAction(ExecuteExit))
	start := createCommand("start", NewAction(ExecuteStart))
	commands = append(commands, show, create, exit, start)
	return NewCommandTree(commands)
}

func createShowCommand() *Command {
	help := createCommand("help", NewAction(ExecuteShowHelp))
	players := createCommand("players", NewAction(ExecuteShowPlayers))
	stats := createCommand("stats", NewAction(ExecuteShowStats))
	return NewCommand("show").
		SubCommand(help).
		SubCommand(players).
		SubCommand(stats)
}

func createCreateCommand() *Command {
	player := createCommand("player", NewAction(ExecuteCreatePlayer))
	return NewCommand("create").
		SubCommand(player)
}

func createCommand(id string, action *Action) *Command {
	return NewCommand(id).WithAction(action)
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
