package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ignaciofalco/new-store/cmd/server/handler"
	"github.com/ignaciofalco/new-store/internal/products"
	"github.com/ignaciofalco/new-store/pkg/store"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}

//go run cmd/server/main.go
