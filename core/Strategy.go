package core

type Strategy interface {
	Handlers() []Handler
	Run(request Request, session Session) Response
}
