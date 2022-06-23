package main

import (
	"fmt"

	config "github.com/Miguel-Florian/KokoChic/Config"
	routers "github.com/Miguel-Florian/KokoChic/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Koko Chic Services")
	fmt.Println("Starting Server ...")

	config.ConnectDB() //run a database
	r := gin.Default()
	r.SetTrustedProxies(nil)

	routers.path(r) //Iniatialize a func router

	r.Run("localhost:8080")
}
