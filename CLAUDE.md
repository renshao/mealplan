# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a full-stack meal planning web application with:
- **Backend**: Go 1.21 with Gin web framework, GORM ORM, and SQLite database
- **Frontend**: Svelte 4 with SvelteKit 2 and Vite 5

## Development Commands

### Backend (Go)
```bash
# Start backend server (from project root)
./start-backend.sh

# Or manually:
cd backend
go run cmd/server/main.go

# Build backend
cd backend
go build -o mealplan-server cmd/server/main.go

# Run Go modules commands
cd backend
go mod tidy
go mod download
```

### Frontend (Svelte/SvelteKit)
```bash
# Start frontend dev server (from project root)
./start-frontend.sh

# Or manually:
cd frontend
npm run dev

# Build frontend for production
cd frontend
npm run build

# Preview production build
cd frontend
npm run preview

# Install dependencies
cd frontend
npm install
```

## Architecture

### Backend Structure
- `cmd/server/main.go` - Application entry point, sets up Gin router with CORS middleware
- `internal/database/` - Database connection and initialization
- `internal/handlers/` - HTTP request handlers for meals and meal plans
- `internal/middleware/` - CORS middleware setup
- `internal/models/` - GORM models (Meal, MealPlan with DayOfWeek/MealType enums)

### Frontend Structure
- `src/lib/api.js` - API client functions for backend communication
- `src/lib/MealInput.svelte` - Meal input component with autocomplete
- `src/routes/+page.svelte` - Main meal planning grid interface
- `src/routes/+layout.svelte` - Layout wrapper

### API Endpoints
- `GET /api/meal-plans` - Get all meal plans
- `GET /api/meal-plans/{dayOfWeek}/{mealType}` - Get specific meal plan
- `POST /api/meal-plans/{dayOfWeek}/{mealType}` - Set meal plan
- `DELETE /api/meal-plans/{dayOfWeek}/{mealType}` - Delete meal plan
- `GET /api/meals/search?name={query}` - Search meals for autocomplete

### Database Schema
- **meals**: id, name (unique), description
- **meal_plans**: id, day_of_week, meal_type, meal_id (FK to meals)

## Key Implementation Notes

- Backend runs on localhost:8080, frontend on localhost:5173
- Database file is `backend/mealplan.db` (SQLite)
- No existing test suite - would need to be implemented if testing is required
- CORS is configured to allow frontend-backend communication
- Meal autocomplete is based on previously entered meal names