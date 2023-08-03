package main

import (
	"fmt"
	"log"

	"github.com/Nickadimas79/jwt-gin/controllers"
	"github.com/Nickadimas79/jwt-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(":::starting app:::")

	fmt.Println("connecting to database")
	models.Connect()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	fmt.Println("starting Gin server")
	err := router.Run(":8080")
	if err != nil {
		log.Println("error starting Gin server:", err)
	}
}
