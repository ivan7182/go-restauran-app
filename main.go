package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=resto_go sslmode=disable"
)

func main() {
	seedDB()
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drinks", getDrinksMenu)
	e.Logger.Fatal(e.Start(":14045"))

}

type MenuType string

const (
	MenuTypeFood   = "food"
	MenuTypeDrinks = "drinks"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int64
	Type      MenuType
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Ayam",
			OrderCode: "ayam",
			Price:     90000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "mie",
			OrderCode: "mie",
			Price:     90000,
			Type:      MenuTypeFood,
		},
	}

	drinksMenu := []MenuItem{
		{
			Name:      "sprite",
			OrderCode: "sprite",
			Price:     90000,
			Type:      MenuTypeDrinks,
		},
		{
			Name:      "fanta",
			OrderCode: "fanta",
			Price:     90000,
			Type:      MenuTypeDrinks,
		},
	}

	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&MenuItem{})
	if err != nil {
		panic(err)
	}

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}

}

func getFoodMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func getDrinksMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuTypeDrinks}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
