package database

import (
	"restaurant-app/internal/model"
	"restaurant-app/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Ayam",
			OrderCode: "ayam",
			Price:     90000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "mie",
			OrderCode: "mie",
			Price:     90000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []model.MenuItem{
		{
			Name:      "sprite",
			OrderCode: "sprite",
			Price:     90000,
			Type:      constant.MenuTypeDrinks,
		},
		{
			Name:      "fanta",
			OrderCode: "fanta",
			Price:     90000,
			Type:      constant.MenuTypeDrinks,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}

}
