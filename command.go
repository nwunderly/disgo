package disgo

import "fmt"

type CommandInfo struct {
	Name string
	Desc string
}

//type CommandBase interface {
//	Invoke(Context) error
//	Info() CommandInfo
//}

type Command struct {
	Name          string
	Desc          string
	QualifiedName string
	Handler       CommandHandler
	Checks        []CommandCheck
	Subcommands   []*Command
	ErrorHandler  ErrorHandler
	PanicHandler  PanicHandler
}

var NilCommand = &Command{
	Handler: func(ctx *Context) error { return nil },
}

func (cmd *Command) Invoke(ctx *Context) error {
	for _, check := range cmd.Checks {
		if !check(ctx) {
			return fmt.Errorf("check for command %s failed", cmd.Name)
		}
	}
	return cmd.Handler(ctx)
}

func (cmd *Command) Info() CommandInfo {
	return CommandInfo{Name: cmd.Name, Desc: cmd.Desc}
}

func (cmd *Command) Check(checks ...CommandCheck) {
	for _, check := range checks {
		cmd.Checks = append(cmd.Checks, check)
	}
}

func (cmd *Command) Subcommand(name string, desc string, callback CommandHandler) (*Command, error) {

	subcommand := &Command{
		Name:    name,
		Desc:    desc,
		Handler: callback,
	}
	err := cmd.AddSubcommand(subcommand)

	return subcommand, err
}

func (cmd *Command) AddSubcommand(subcommand *Command) error {
	_, commandExists := cmd.GetSubcommand(subcommand.Name)
	if commandExists {
		return fmt.Errorf("subcommand %s already exists", subcommand.Name)
	}
	subcommand.QualifiedName = cmd.QualifiedName + subcommand.Name
	cmd.Subcommands = append(cmd.Subcommands, subcommand)
	return nil
}

func (cmd *Command) GetSubcommand(name string) (*Command, bool) {
	for _, subcommand := range cmd.Subcommands {
		if subcommand == nil {
			continue
		}
		if subcommand.Name == name {
			return subcommand, true
		}
	}
	return NilCommand, false
}

func (cmd *Command) RemoveSubcommand(name string) {
	for i, subcommand := range cmd.Subcommands {
		if subcommand.Name == name {
			cmd.Subcommands = deleteCommand(cmd.Subcommands, i)
		}
	}
}
