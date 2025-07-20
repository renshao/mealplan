package main

import (
	"log"
	"mealplan-backend/internal/database"
	"mealplan-backend/internal/handlers"
	"mealplan-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	router.Use(middleware.SetupCORS())

	api := router.Group("/api")
	{
		meals := api.Group("/meals")
		{
			meals.GET("", handlers.GetAllMeals)
			meals.GET("/:id", handlers.GetMealByID)
			meals.POST("", handlers.CreateMeal)
			meals.PUT("/:id", handlers.UpdateMeal)
			meals.DELETE("/:id", handlers.DeleteMeal)
			meals.GET("/search", handlers.SearchMeals)
		}

		mealPlans := api.Group("/meal-plans")
		{
			mealPlans.GET("", handlers.GetAllMealPlans)
			mealPlans.GET("/:dayOfWeek/:mealType", handlers.GetMealPlan)
			mealPlans.POST("/:dayOfWeek/:mealType", handlers.SetMealPlan)
			mealPlans.DELETE("/:dayOfWeek/:mealType", handlers.DeleteMealPlan)
		}
	}

	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}