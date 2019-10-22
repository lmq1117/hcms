package controller

import (
	"context"
	"github.com/lmq1117/hcms/service"
)

type OrderController struct {
	Ctx     context.Context
	Service service.OrderService
}
