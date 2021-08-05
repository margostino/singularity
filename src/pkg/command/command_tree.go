package command

type CommandTree struct {
	commands []*Command
}

func NewCommandTree(commands []*Command) *CommandTree {
	return &CommandTree{commands}
}

func (tree *CommandTree) IsValidPlan(plan []string) bool {
	for _, command := range tree.commands {
		if command.Validate(plan) {
			return true
		}
	}
	return false
}

func (tree *CommandTree) LookupCommand(plan []string) *Command {
	var result *Command = nil
	for _, command := range tree.commands {
		result = command.GetLastCommand(plan)
		if result != nil {
			break
		}
	}
	return result
}
