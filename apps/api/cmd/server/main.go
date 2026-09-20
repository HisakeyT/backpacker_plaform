package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/HisakeyT/backpacker-platform/internal/auth"
	"github.com/HisakeyT/backpacker-platform/internal/database"
	"github.com/HisakeyT/backpacker-platform/internal/travel"
	"github.com/HisakeyT/backpacker-platform/internal/travel_plan"
	"github.com/HisakeyT/backpacker-platform/internal/user"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtManager := auth.NewJWTManager(jwtSecret)

	userRepository := user.NewGormRepository(db)
	userUseCase := user.NewUseCase(userRepository, jwtManager)
	userHandler := user.NewHandler(userUseCase)

	travelRepository := travel.NewGormRepository(db)
	travelUseCase := travel.NewUseCase(travelRepository)
	travelHandler := travel.NewHandler(travelUseCase)

	travelPlanRepository := travel_plan.NewGormRepository(db)
	travelPlanUseCase := travel_plan.NewUseCase(travelRepository, travelPlanRepository)
	travelPlanHandler := travel_plan.NewHandler(travelPlanUseCase)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	protected := r.Group("/users")
	protected.Use(auth.AuthMiddleware(jwtManager))
	{
		protected.GET("/me", userHandler.Me)
		protected.PATCH("me", userHandler.UpdateMe)
	}

	travels := protected.Group("/travels")
	{
		travels.POST("", travelHandler.CreateTravel)
		travels.POST("/:travel_id/plans", travelPlanHandler.CreateTravelPlan)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
