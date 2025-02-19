package resto

import "restaurant-app/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
