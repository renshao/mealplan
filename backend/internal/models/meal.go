package models

type Meal struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}

type DayOfWeek string

const (
	Monday    DayOfWeek = "MONDAY"
	Tuesday   DayOfWeek = "TUESDAY"
	Wednesday DayOfWeek = "WEDNESDAY"
	Thursday  DayOfWeek = "THURSDAY"
	Friday    DayOfWeek = "FRIDAY"
	Saturday  DayOfWeek = "SATURDAY"
	Sunday    DayOfWeek = "SUNDAY"
)

type MealType string

const (
	Breakfast MealType = "BREAKFAST"
	Lunch     MealType = "LUNCH"
	Dinner    MealType = "DINNER"
)

type MealPlan struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	DayOfWeek DayOfWeek `json:"dayOfWeek" gorm:"not null"`
	MealType  MealType  `json:"mealType" gorm:"not null"`
	MealID    uint      `json:"mealId" gorm:"not null"`
	Meal      Meal      `json:"meal" gorm:"foreignKey:MealID"`
}

func (MealPlan) TableName() string {
	return "meal_plans"
}

func (Meal) TableName() string {
	return "meals"
}