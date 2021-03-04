package disgo

type CommandHandler func(*Context) error

type CommandCheck func(*Context) bool

type ErrorHandler func(*Context, error)

type PanicHandler func(*Context, interface{})
