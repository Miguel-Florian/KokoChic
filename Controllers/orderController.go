package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Order(c *gin.Context) {
	c.HTML(http.StatusOK, "order.html", gin.H{
		"title": "Order",
	})
}
