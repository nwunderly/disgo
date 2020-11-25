package commands

type commandInvoke func(Context) error

type Command struct {
	Name   string
	Invoke commandInvoke
}
