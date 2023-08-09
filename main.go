package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yueyue000/webook/internal/web"
)

func main() {
	server := gin.Default()
	web.RegisterRoutes(server)
	server.Run(":8080")
}
