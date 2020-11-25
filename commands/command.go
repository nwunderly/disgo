package commands

type commandInvokeCallback func(Context) error
type CommandInfo struct {
	Name string
	Desc string
}

type CommandBase interface {
	Invoke(Context) error
	Info() CommandInfo
}

type Command struct {
	Name           string
	Desc           string
	InvokeCallback commandInvokeCallback
}

func (cmd Command) Invoke(ctx Context) error {
	return cmd.InvokeCallback(ctx)
}

func (cmd Command) Info() CommandInfo {
	return CommandInfo{Name: cmd.Name, Desc: cmd.Desc}
}
