package menu

import (
	"restaurant-app/internal/model"
)

type Repository interface {
	GetMenu(MenuType string) ([]model.MenuItem, error)
}
