package command

import (
	"fmt"
	"org.gene/singularity/pkg/action"
	"org.gene/singularity/pkg/config"
	"testing"
)

func ExecuteDummyAction() {
	fmt.Println("dummy action")
}

func GetDummyAction() *action.Action {
	return &action.Action{
		Function: ExecuteDummyAction,
	}
}

func assertTrue(plan []string, commandMap *CommandMap, t *testing.T) {
	if !commandMap.IsValidPlan(plan) {
		t.Fatalf(`Command %q for commands is not valid`, plan)
	}
}

func assertFalse(plan []string, commandMap *CommandMap, t *testing.T) {
	if commandMap.IsValidPlan(plan) {
		t.Fatalf(`Command %q for commands is not valid`, plan)
	}
}

func TestLoadCommandMap(t *testing.T) {
	config.LoadCommandsConfiguration()
	commandMap := Load()

	plan := []string{"help"}
	assertTrue(plan, commandMap, t)

	plan = []string{"show", "players"}
	assertTrue(plan, commandMap, t)

	plan = []string{"exit"}
	assertTrue(plan, commandMap, t)

	plan = []string{"show", "stats"}
	assertTrue(plan, commandMap, t)

	plan = []string{"start"}
	assertTrue(plan, commandMap, t)

	plan = []string{"deactivate"}
	assertTrue(plan, commandMap, t)

	plan = []string{"create", "player"}
	assertTrue(plan, commandMap, t)

	plan = []string{"select", "player", "one"}
	assertTrue(plan, commandMap, t)
}

func TestValidCommandMapWithMultipleOptions(t *testing.T) {
	level2a := NewCommand("help").WithAction(GetDummyAction())
	level2b := NewCommand("player").WithAction(GetDummyAction())
	root := NewCommand("show").SubCommand(level2a).SubCommand(level2b)
	commands := map[string]*Command{"show": root}
	commandMap := NewCommandMap(commands)

	plan := []string{"show", "help"}
	assertTrue(plan, commandMap, t)

	plan = []string{"show", "player"}
	assertTrue(plan, commandMap, t)

	plan = []string{"show", "player", "invalid"}
	assertFalse(plan, commandMap, t)
}

func TestValidCommandMapWith2Levels(t *testing.T) {
	level2 := NewCommand("command2_2").WithAction(GetDummyAction())
	root := NewCommand("command1_2").SubCommand(level2)
	commands := map[string]*Command{"command1_2": root}
	commandMap := NewCommandMap(commands)
	plan := []string{"command1_2", "command2_2"}
	assertTrue(plan, commandMap, t)
}

func TestValidCommandMapWith3Levels(t *testing.T) {
	level3 := NewCommand("command3_3").WithAction(GetDummyAction())
	level2 := NewCommand("command2_3").SubCommand(level3)
	root := NewCommand("command1_3").SubCommand(level2)
	commands := map[string]*Command{"command1_3": root}
	commandMap := NewCommandMap(commands)
	plan := []string{"command1_3", "command2_3", "command3_3"}
	assertTrue(plan, commandMap, t)
}

func TestInvalidCommandMap(t *testing.T) {
	level3 := NewCommand("command3_3").WithAction(GetDummyAction())
	level2 := NewCommand("command2_3").SubCommand(level3)
	root := NewCommand("command1_3").SubCommand(level2)
	commands := map[string]*Command{"command1_3": root}
	commandMap := NewCommandMap(commands)

	plan := []string{"command1_3", "command2_3"}
	assertFalse(plan, commandMap, t)

	level2 = NewCommand("command2_2").WithAction(GetDummyAction())
	root = NewCommand("command1_2").SubCommand(level2)
	commands = map[string]*Command{"command1_2": root}
	commandMap = NewCommandMap(commands)

	plan = []string{"command1_3", "command2_3", "command3_3"}
	assertFalse(plan, commandMap, t)

	level2 = NewCommand("command2_2").WithAction(GetDummyAction())
	root = NewCommand("command1_2").SubCommand(level2)
	commands = map[string]*Command{"command1_2": root}
	commandMap = NewCommandMap(commands)

	plan = []string{"command1_2", "command2_2", "extra"}
	assertFalse(plan, commandMap, t)

	level2 = NewCommand("command2_2").WithAction(GetDummyAction())
	root = NewCommand("command1_2").SubCommand(level2)
	commands = map[string]*Command{"command1_2": root}
	commandMap = NewCommandMap(commands)

	plan = []string{"invalid"}
	assertFalse(plan, commandMap, t)

	level2 = NewCommand("command2_2").WithAction(GetDummyAction())
	root = NewCommand("command1_2").SubCommand(level2)
	commands = map[string]*Command{"command1_2": root}
	commandMap = NewCommandMap(commands)

	plan = []string{"invalid", "invalid"}
	assertFalse(plan, commandMap, t)
}
