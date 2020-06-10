package core

type Handler interface {
	Handle(request Request, session Session)
}
