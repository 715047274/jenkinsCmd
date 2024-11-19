package application

import (
	"fmt"
	"github.com/gin/demo/internal/domain"
)

type AnalyticService struct {
}

func (analytic *AnalyticService) HandleCarEvent(event interface{}) {
	if accountEvent, ok := event.(domain.AccountEvent); ok {
		fmt.Printf("Running analytic service event handle for UserID=%s, ProductID=%d, Quantity=%s\n",
			accountEvent.Owner, accountEvent.Balance, accountEvent.Currency)
	}
}
