package resto

import (
	"restaurant-app/internal/model"
	"restaurant-app/internal/model/constant"
	"restaurant-app/internal/repository/menu"
	"restaurant-app/internal/repository/order"

	"github.com/google/uuid"
)

type restoUseCase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
}

func GetUseCase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUseCase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
	}
}

func (r *restoUseCase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUseCase) Order(request model.OrderMenuRequest) (model.Order, error) {
	ProductOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		ProductOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData[0].OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menuData[0].Price * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProcessed,
		ProductOrders: ProductOrderData,
	}

	createOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createOrderData, nil
}

func (r *restoUseCase) getOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderID)
	if err != nil {
		return orderData, err
	}
	return orderData, nil
}
