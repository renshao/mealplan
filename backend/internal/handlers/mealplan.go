package handlers

import (
	"mealplan-backend/internal/database"
	"mealplan-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SetMealPlanRequest struct {
	MealName string `json:"mealName" binding:"required"`
}

func GetAllMealPlans(c *gin.Context) {
	var mealPlans []models.MealPlan
	database.DB.Preload("Meal").Find(&mealPlans)
	c.JSON(http.StatusOK, mealPlans)
}

func GetMealPlan(c *gin.Context) {
	dayOfWeek := models.DayOfWeek(c.Param("dayOfWeek"))
	mealType := models.MealType(c.Param("mealType"))

	var mealPlan models.MealPlan
	if err := database.DB.Preload("Meal").Where("day_of_week = ? AND meal_type = ?", dayOfWeek, mealType).First(&mealPlan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal plan not found"})
		return
	}

	c.JSON(http.StatusOK, mealPlan)
}

func SetMealPlan(c *gin.Context) {
	dayOfWeek := models.DayOfWeek(c.Param("dayOfWeek"))
	mealType := models.MealType(c.Param("mealType"))

	var req SetMealPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var meal models.Meal
	if err := database.DB.Where("name = ?", req.MealName).First(&meal).Error; err != nil {
		meal = models.Meal{
			Name:        req.MealName,
			Description: "",
		}
		database.DB.Create(&meal)
	}

	var existingPlan models.MealPlan
	if err := database.DB.Where("day_of_week = ? AND meal_type = ?", dayOfWeek, mealType).First(&existingPlan).Error; err == nil {
		existingPlan.MealID = meal.ID
		database.DB.Save(&existingPlan)
		database.DB.Preload("Meal").First(&existingPlan, existingPlan.ID)
		c.JSON(http.StatusOK, existingPlan)
		return
	}

	newPlan := models.MealPlan{
		DayOfWeek: dayOfWeek,
		MealType:  mealType,
		MealID:    meal.ID,
	}
	database.DB.Create(&newPlan)
	database.DB.Preload("Meal").First(&newPlan, newPlan.ID)
	c.JSON(http.StatusCreated, newPlan)
}

func DeleteMealPlan(c *gin.Context) {
	dayOfWeek := models.DayOfWeek(c.Param("dayOfWeek"))
	mealType := models.MealType(c.Param("mealType"))

	var mealPlan models.MealPlan
	if err := database.DB.Where("day_of_week = ? AND meal_type = ?", dayOfWeek, mealType).First(&mealPlan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal plan not found"})
		return
	}

	database.DB.Delete(&mealPlan)
	c.JSON(http.StatusOK, gin.H{"message": "Meal plan deleted successfully"})
}