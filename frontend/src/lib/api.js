const API_BASE = 'http://localhost:8080/api';

export async function getMealPlans() {
  const response = await fetch(`${API_BASE}/meal-plans`);
  return response.json();
}

export async function setMealPlan(dayOfWeek, mealType, mealName) {
  const response = await fetch(`${API_BASE}/meal-plans/${dayOfWeek}/${mealType}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ mealName })
  });
  return response.json();
}

export async function deleteMealPlan(dayOfWeek, mealType) {
  const response = await fetch(`${API_BASE}/meal-plans/${dayOfWeek}/${mealType}`, {
    method: 'DELETE'
  });
  return response.ok;
}

export async function searchMeals(query) {
  const response = await fetch(`${API_BASE}/meals/search?name=${encodeURIComponent(query)}`);
  return response.json();
}