package main

import (
	"fmt"

	config "github.com/Miguel-Florian/KokoChic/Config"
	controllers "github.com/Miguel-Florian/KokoChic/Controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Koko Chic Services")
	fmt.Println("Starting Server ...")

	config.ConnectDB() //run a database
	r := gin.Default()

	//Iniatialize a func router
	router := r.Group("/kokochic")
	{
		router.GET("/", controllers.Index())
		router.GET("/about", controllers.About())
		router.GET("/contact", controllers.Contact())
	}

	r.Run("localhost:8080")
}
