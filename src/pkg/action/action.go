package action

type Action struct {
	Function      func()
	InputFunction func([]string)
}

func NewAction(function func()) *Action {
	return &Action{
		Function: function,
	}
}

func NewInputAction(function func([]string)) *Action {
	return &Action{
		InputFunction: function,
	}
}

func (a Action) Execute() {
	a.Function()
}
