# Meal Planning Web App

A full-stack meal planning application built with Svelte frontend and Golang Gin backend.

## Features

- Weekly meal planning grid (breakfast, lunch, dinner for each day)
- Autocomplete functionality for meal entries
- Persistent storage of meal plans
- Automatic meal suggestions based on previous entries

## Technology Stack

**Backend:**
- Go 1.21
- Gin Web Framework
- GORM (Go ORM)
- SQLite Database

**Frontend:**
- Svelte 4
- SvelteKit 2
- Vite 5

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Node.js 18+ and npm

### Running the Application

1. **Start the Go Backend:**
   ```bash
   ./start-backend-go.sh
   ```
   Or manually:
   ```bash
   cd backend
   go run cmd/server/main.go
   ```
   The backend will run on http://localhost:8080

   **Alternative Java Backend:**
   ```bash
   ./start-backend.sh
   ```
   (Java backend is in the `backend-java` folder)

2. **Start the Frontend:**
   ```bash
   ./start-frontend.sh
   ```
   Or manually:
   ```bash
   cd frontend
   npm run dev
   ```
   The frontend will run on http://localhost:5173

3. **Access the Application:**
   Open http://localhost:5173 in your browser

## API Endpoints

- `GET /api/meal-plans` - Get all meal plans
- `GET /api/meal-plans/{dayOfWeek}/{mealType}` - Get specific meal plan
- `POST /api/meal-plans/{dayOfWeek}/{mealType}` - Set meal plan
- `DELETE /api/meal-plans/{dayOfWeek}/{mealType}` - Delete meal plan
- `GET /api/meals/search?name={query}` - Search meals for autocomplete

## Usage

1. Click on any cell in the weekly grid
2. Type a meal name - autocomplete suggestions will appear based on previous entries
3. Press Enter or click outside to save
4. Previously entered meals will be available as autocomplete options

## Database

**Go Backend:** Uses SQLite database (`mealplan.db` file). Data persists between restarts.
