package commands

type HandlerType func(args []string) error

type Command struct {
	Handler     HandlerType
	Description string
}
