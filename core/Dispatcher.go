package core

type Dispatcher interface {
	Dispatch(request Request) Strategy
}
