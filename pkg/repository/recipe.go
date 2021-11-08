package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type RecipeRepository interface {
	FindAll() ([]*domain.Recipe, error)
	FindByUUID(uuid string) (*domain.Recipe, error)
}

type Recipe struct {
	db *sqlx.DB
}

func NewRecipeRepository(db *sqlx.DB) RecipeRepository {
	return &Recipe{db: db}
}

func (r *Recipe) FindAll() ([]*domain.Recipe, error) {
	var recipes []*domain.Recipe

	rows, err := r.db.Queryx("SELECT uuid, name, description, timing FROM recipes")
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

func (r *Recipe) FindByUUID(uuid string) (*domain.Recipe, error) {
	var recipe domain.Recipe

	err := r.db.Get(&recipe, `
		SELECT
			r.uuid,
			r.name,
			r.description,
			r.timing,
			s.uuid,
			s.index,
			s.instruction
		FROM recipes r
		INNER JOIN steps s ON r.uuid = s.recipe_id
		WHERE r.uuid = $1
	`, uuid)
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}
