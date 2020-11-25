package commands

type commandInvokeCallback func(Context) error

type CommandBase interface {
	Invoke(Context) error
}

type Command struct {
	Name   string
	InvokeCallback commandInvokeCallback
}

func (cmd Command) Invoke(ctx Context) error {
	return cmd.InvokeCallback(ctx)
}