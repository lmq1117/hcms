package controller

import (
	"context"
	"hcms/service"
)

type OrderController struct {
	Ctx     context.Context
	Service service.OrderService
}
