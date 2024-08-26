package controllers

import (
	"LearnEcho/models"
	"context"
	"log"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StartProductsController struct {
	Db *mongo.Database
}

func (s *StartProductsController) GetProducts(ctx echo.Context) error {
	cl := s.Db.Collection("products")
	var products []bson.M
	cursor, err := cl.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.Background(), &products)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(200, products)
}

func (s *StartProductsController) CreatedProduct(ctx echo.Context) error {
	product := &models.Products{}
	cl := s.Db.Collection("products")
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(400, bson.M{
			"error": "Invalid product data",
		})
	}

	_, err := cl.InsertOne(ctx.Request().Context(), product)
	if err != nil {
		return ctx.JSON(500, bson.M{
			"error": "Failed to create product",
		})
	}

	return ctx.JSON(201, bson.M{
		"msg": product,
	})
}
