package action

type Action struct {
	Apply func()
}

func NewAction(apply func()) *Action {
	return &Action{
		Apply: apply,
	}
}

func (a Action) Execute() {
	a.Apply()
}
