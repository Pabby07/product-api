package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/product-api/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/products", handlers.CreateProduct)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
