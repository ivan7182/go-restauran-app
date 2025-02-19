package main

import (
	"restaurant-app/internal/database"
	"restaurant-app/internal/delivery/rest"
	mRepo "restaurant-app/internal/repository/menu"
	rUsecase "restaurant-app/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=resto_go sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUseCase(menuRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))

}
