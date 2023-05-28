package main

import (
	"architecture-test/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler.RegistRoute(router)
	// サーバーを起動
	router.Run(":8080")
}
