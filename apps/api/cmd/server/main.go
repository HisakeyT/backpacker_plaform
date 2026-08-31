package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/HisakeyT/backpacker-platform/internal/database"
	"github.com/HisakeyT/backpacker-platform/internal/user"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully")

	userRepository := user.NewGormRepository(db)
	userUseCase := user.NewUseCase(userRepository)
	userHandler := user.NewHandler(userUseCase)

	_ = userUseCase

	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
