package constant

import "restaurant-app/internal/model"

const (
	OrderStatusProcessed model.OrderStatus = "processed"
	OrderStatusFinised   model.OrderStatus = "finised"
	OrderStatusFailed    model.OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing model.ProductOrderStatus = "Preparing"
	ProductOrderStatusFinised   model.ProductOrderStatus = "finished"
)
