package core

type BizService interface {
	Service(request Request) Response
}
