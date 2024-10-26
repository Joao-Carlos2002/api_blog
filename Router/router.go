package router

import (
	"github.com/Joao-Carlos2002/api_blog/Router/routes"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/login", routes.LoginPost)
	router.Run(":8080")
}
