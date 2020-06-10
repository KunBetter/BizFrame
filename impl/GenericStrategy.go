package impl

import (
	"github.com/KunBetter/BizFrame/core"
)

type GenericStrategy struct {
	handlers []core.Handler
}

func (gs *GenericStrategy) Handlers() []core.Handler {
	return gs.handlers
}

func (gs *GenericStrategy) Run(request core.Request, session core.Session) core.Response {
	for _, handler := range gs.handlers {
		handler.Handle(request, session)
	}

	return nil
}
