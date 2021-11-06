package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RecipeRepository interface {
	FindAll() (string, error)
}

type Recipe struct {
	db *pgxpool.Pool
}

func NewRecipeRepository(db *pgxpool.Pool) RecipeRepository {
	return &Recipe{db: db}
}

func (r *Recipe) FindAll() (string, error) {
	// var recipes []*domain.Recipe

	var name string
	err := r.db.QueryRow(context.Background(), "SELECT name FROM recipes LIMIT 1").Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}
