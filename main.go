package main

import (
	"fmt"

	//config "github.com/Miguel-Florian/KokoChic/Config"
	controllers "github.com/Miguel-Florian/KokoChic/Controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Koko Chic Services")
	fmt.Println("Starting Server ...")

	//config.ConnectDB() //run a database
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.LoadHTMLGlob("templates/*")

	//Iniatialize a func router
	router := r.Group("/kokochic")
	{
		router.GET("/home", controllers.Index)
	}
	r.Run("localhost:8080")
}
