package impl

import (
	"github.com/KunBetter/BizFrame/core"
)

type GenericBizService struct {
	dispatcher core.Dispatcher
}

func (gb *GenericBizService) Service(request core.Request) core.Response {
	strategy := gb.dispatcher.Dispatch(request)
	return strategy.Run(request, nil)
}
