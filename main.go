package main

import (
	"LearnEcho/configs"
	"LearnEcho/controllers"
	"LearnEcho/models"

	"github.com/labstack/echo"
)

func main() {

	app := echo.New()
	Db := configs.ConnectionDatabase()
	models.CreateCollectionProducts(Db)
	productsController := controllers.StartProductsController{Db}

	app.GET("/api/v1/products", productsController.GetProducts)
	app.POST("/api/v1/products", productsController.CreatedProduct)
	app.Start("localhost:2000")

}
