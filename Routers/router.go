package routers

import (
	controllers "github.com/Miguel-Florian/KokoChic/Controllers"
	"github.com/gin-gonic/gin"
)

func path(router *gin.Engine) {
	r := router.Group("/kokochic")
	{
		r.GET("/home", controllers.Index)
		r.GET("/about", controllers.Order)
	}
}
