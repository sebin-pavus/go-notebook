package main

import (
	"go-notebook/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	web.NewServer(router)

	router.Run("0.0.0.0:8080")
}
