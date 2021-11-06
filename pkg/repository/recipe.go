package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type RecipeRepository interface {
	FindAll() ([]*domain.Recipe, error)
}

type Recipe struct {
	db *sqlx.DB
}

func NewRecipeRepository(db *sqlx.DB) RecipeRepository {
	return &Recipe{db: db}
}

func (r *Recipe) FindAll() ([]*domain.Recipe, error) {
	var recipes []*domain.Recipe

	rows, err := r.db.Queryx("SELECT uuid, name FROM recipes")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var recipe domain.Recipe
		err = rows.StructScan(&recipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}
