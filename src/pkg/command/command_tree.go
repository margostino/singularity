package command

import "org.gene/singularity/pkg/action"

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

func (tree *CommandTree) LookupAction(plan []string) *action.Action {
	var action *action.Action = nil
	for _, command := range tree.commands {
		action = command.GetAction(plan)
		if action != nil {
			break
		}
	}
	return action
}
