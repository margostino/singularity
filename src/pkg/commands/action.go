package commands

type Action struct {
	apply func()
	exit  bool
}

func NewAction(apply func()) *Action {
	return &Action{
		apply: apply,
	}
}

func (a Action) Execute() {
	a.apply()
}
