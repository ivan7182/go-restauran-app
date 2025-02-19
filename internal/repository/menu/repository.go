package menu

import (
	"restaurant-app/internal/model"
)

type Repository interface {
	GetMenuList(MenuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) ([]model.MenuItem, error)
}
