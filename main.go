package main

import (
	"fmt"
	"github.com/Nickadimas79/jwt-gin/models"
	"log"

	"github.com/Nickadimas79/jwt-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(":::starting app:::")

	fmt.Println("connecting to database")
	models.Connect()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)

	fmt.Println("starting Gin server")
	err := router.Run(":8080")
	if err != nil {
		log.Println("error starting Gin server:", err)
	}
}
