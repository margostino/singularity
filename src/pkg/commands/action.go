package commands

type Action struct {
	name  string
	apply func()
	exit bool
}

func (a Action) Execute() {
	a.apply()
}
