package main

import (
	"github.com/gin-gonic/gin"
	"gomentoring/repository"
	"gomentoring/handlers"
	"gomentoring/usecase"
)

func main() {
	r := gin.Default()

	repo := repository.NewInMemoryProductRepo()
	uc := usecase.NewProductUsecase(repo)
	handlers.NewProductHandler(r, uc)

	r.Run(":8080")
}


