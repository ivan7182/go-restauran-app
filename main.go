package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

type MenuDrink struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "bakmi",
			OrderCode: "bakmi",
			Price:     20000,
		},
		{
			Name:      "ayam",
			OrderCode: "ayam",
			Price:     23000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getFoodMenuDrink(c echo.Context) error {
	DrinkMenu := []MenuDrink{
		{
			Name:      "Sprite",
			OrderCode: "Sprite",
			Price:     10000,
		},
		{
			Name:      "Fanta",
			OrderCode: "Fanta",
			Price:     21000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": DrinkMenu,
	})
}

func main() {
	e := echo.New()
	//localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drinks", getFoodMenuDrink)
	e.Logger.Fatal(e.Start(":14045"))
}
