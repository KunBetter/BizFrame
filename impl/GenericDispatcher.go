package impl

import (
	"github.com/KunBetter/BizFrame/core"
)

type GenericDispatcher struct {
	StrategyMap map[string]core.Strategy
}

func (gd *GenericDispatcher) Dispatch(request core.Request) core.Strategy {
	bizName := request.BizName()
	if strategy, ok := gd.StrategyMap[bizName]; ok {
		return strategy
	}

	return nil
}
