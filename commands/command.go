package commands

type CommandHandler func(*Context) error
type CommandInfo struct {
	Name string
	Desc string
}

//type CommandBase interface {
//	Invoke(Context) error
//	Info() CommandInfo
//}

type Command struct {
	Name    string
	Desc    string
	Handler CommandHandler
}


var NilCommand = &Command{
	Handler: func(ctx *Context) error {return nil},
}


func (cmd *Command) Invoke(ctx *Context) error {
	return cmd.Handler(ctx)
}

func (cmd *Command) Info() CommandInfo {
	return CommandInfo{Name: cmd.Name, Desc: cmd.Desc}
}
