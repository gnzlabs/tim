package command

type Handler interface {
	Name() string
	Handle(Message) error
}
