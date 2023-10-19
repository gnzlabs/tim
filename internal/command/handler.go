package command

type Handler interface {
	Name() string
	Handle(string) error
}
