package handlers

import (
	"mealplan-backend/internal/database"
	"mealplan-backend/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllMeals(c *gin.Context) {
	var meals []models.Meal
	database.DB.Find(&meals)
	c.JSON(http.StatusOK, meals)
}

func GetMealByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var meal models.Meal
	if err := database.DB.First(&meal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	c.JSON(http.StatusOK, meal)
}

func CreateMeal(c *gin.Context) {
	var meal models.Meal
	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingMeal models.Meal
	if err := database.DB.Where("name = ?", meal.Name).First(&existingMeal).Error; err == nil {
		c.JSON(http.StatusOK, existingMeal)
		return
	}

	database.DB.Create(&meal)
	c.JSON(http.StatusCreated, meal)
}

func UpdateMeal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var meal models.Meal
	if err := database.DB.First(&meal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	var updateData models.Meal
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&meal).Updates(updateData)
	c.JSON(http.StatusOK, meal)
}

func DeleteMeal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DB.Delete(&models.Meal{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})
}

func SearchMeals(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
		return
	}

	var meals []models.Meal
	searchPattern := "%" + strings.ToLower(name) + "%"
	database.DB.Where("LOWER(name) LIKE ?", searchPattern).Find(&meals)
	c.JSON(http.StatusOK, meals)
}