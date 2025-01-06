package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	dbConnection, err := db.ConnectDb()
	if err != nil { // TODO: handle error condition {
		panic(err)
	}

	// Product
	productRepository := repository.NewProductRepository(dbConnection)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	//server.PUT("/products/:id", productController.UpdateProduct)
	//server.DELETE("/products/:id", productController.DeleteProduct)
	//server.GET("/products/:id", productController.GetProductByName)

	// Ping
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	err = server.Run(":8000")
	if err != nil { // TODO: handle error condition {
		panic(err)
	}
}
