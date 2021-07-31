package commands

import (
	"fmt"
	"testing"
)

type DummyAction struct {
	Action
}

func (a *DummyAction) Execute() {
	fmt.Println("dummy action")
}

func GetDummyAction() *DummyAction {
	dummyAction := DummyAction{}
	action := Action{
		name: "show rules",
	}
	action.apply = (dummyAction).Execute
	return &DummyAction{action}
}

func assertTrue(plan []string, commandTree *CommandTree, t *testing.T) {
	if !commandTree.IsValidPlan(plan) {
		t.Fatalf(`Command %q for commands is not valid`, plan)
	}
}

func assertFalse(plan []string, commandTree *CommandTree, t *testing.T) {
	if commandTree.IsValidPlan(plan) {
		t.Fatalf(`Command %q for commands is not valid`, plan)
	}
}

func TestLoadCommandTree(t *testing.T) {
	commandTree := Load()
	plan := []string{"show", "help"}
	assertTrue(plan, commandTree, t)
	plan = []string{"show", "players"}
	assertTrue(plan, commandTree, t)
	plan = []string{"exit"}
	assertTrue(plan, commandTree, t)
}

func TestValidCommandTreeWithMultipleOptions(t *testing.T) {
	commands := make([]*Command, 0)
	level2a := NewCommand("help").WithAction(&GetDummyAction().Action)
	level2b := NewCommand("player").WithAction(&GetDummyAction().Action)
	root := NewCommand("show").SubCommand(level2a).SubCommand(level2b)
	commands = append(commands, root)
	commandTree := NewCommandTree(commands)
	plan := []string{"show", "help"}
	assertTrue(plan, commandTree, t)

	plan = []string{"show", "player"}
	assertTrue(plan, commandTree, t)

	plan = []string{"show", "player", "invalid"}
	assertFalse(plan, commandTree, t)
}

func TestValidCommandTreeWith2Levels(t *testing.T) {
	commands := make([]*Command, 0)
	level2 := NewCommand("command2_2").WithAction(&GetDummyAction().Action)
	root := NewCommand("command1_2").SubCommand(level2)
	commands = append(commands, root)
	commandTree := NewCommandTree(commands)
	plan := []string{"command1_2", "command2_2"}
	assertTrue(plan, commandTree, t)
}

func TestValidCommandTreeWith3Levels(t *testing.T) {
	commands := make([]*Command, 0)
	level3 := NewCommand("command3_3").WithAction(&GetDummyAction().Action)
	level2 := NewCommand("command2_3").SubCommand(level3)
	root := NewCommand("command1_3").SubCommand(level2)
	commands = append(commands, root)
	commandTree := NewCommandTree(commands)
	plan := []string{"command1_3", "command2_3", "command3_3"}
	assertTrue(plan, commandTree, t)
}

func TestInvalidCommandTree(t *testing.T) {
	commands := make([]*Command, 0)
	level3 := NewCommand("command3_3").WithAction(&GetDummyAction().Action)
	level2 := NewCommand("command2_3").SubCommand(level3)
	root := NewCommand("command1_3").SubCommand(level2)
	commands = append(commands, root)
	commandTree := NewCommandTree(commands)
	plan := []string{"command1_3", "command2_3"}
	assertFalse(plan, commandTree, t)

	commands = make([]*Command, 0)
	level2 = NewCommand("command2_2").WithAction(&GetDummyAction().Action)
	root = NewCommand("command1_2").SubCommand(level2)
	commands = append(commands, root)
	commandTree = NewCommandTree(commands)
	plan = []string{"command1_3", "command2_3", "command3_3"}
	assertFalse(plan, commandTree, t)

	commands = make([]*Command, 0)
	level2 = NewCommand("command2_2").WithAction(&GetDummyAction().Action)
	root = NewCommand("command1_2").SubCommand(level2)
	commands = append(commands, root)
	commandTree = NewCommandTree(commands)
	plan = []string{"command1_2", "command2_2", "extra"}
	assertFalse(plan, commandTree, t)

	commands = make([]*Command, 0)
	level2 = NewCommand("command2_2").WithAction(&GetDummyAction().Action)
	root = NewCommand("command1_2").SubCommand(level2)
	commands = append(commands, root)
	commandTree = NewCommandTree(commands)
	plan = []string{"invalid"}
	assertFalse(plan, commandTree, t)

	commands = make([]*Command, 0)
	level2 = NewCommand("command2_2").WithAction(&GetDummyAction().Action)
	root = NewCommand("command1_2").SubCommand(level2)
	commands = append(commands, root)
	commandTree = NewCommandTree(commands)
	plan = []string{"invalid", "invalid"}
	assertFalse(plan, commandTree, t)
}
